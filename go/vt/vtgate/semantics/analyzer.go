/*
Copyright 2020 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package semantics

import (
	"fmt"
	"runtime/debug"
	"strings"

	querypb "vitess.io/vitess/go/vt/proto/query"
	vtrpcpb "vitess.io/vitess/go/vt/proto/vtrpc"
	"vitess.io/vitess/go/vt/sqlparser"
	"vitess.io/vitess/go/vt/vterrors"
)

type (
	// analyzer is a struct to work with analyzing the query.
	analyzer struct {
		scoper *scoper
		tables *tableCollector
		binder *binder

		exprTypes    map[sqlparser.Expr]querypb.Type
		err          error
		inProjection []bool

		projErr error
	}
)

// newAnalyzer create the semantic analyzer
func newAnalyzer(dbName string, si SchemaInformation) *analyzer {
	s := newScoper()
	return &analyzer{
		exprTypes: map[sqlparser.Expr]querypb.Type{},
		scoper:    s,
		tables:    newTableCollector(s, si, dbName),
		binder:    newBinder(),
	}
}

// Analyze analyzes the parsed query.
func Analyze(statement sqlparser.SelectStatement, currentDb string, si SchemaInformation) (*SemTable, error) {
	analyzer := newAnalyzer(currentDb, si)
	// Initial scope
	err := analyzer.analyze(statement)
	if err != nil {
		return nil, err
	}
	return &SemTable{
		ExprBaseTableDeps: analyzer.binder.exprRecursiveDeps,
		ExprDeps:          analyzer.binder.exprDeps,
		exprTypes:         analyzer.exprTypes,
		Tables:            analyzer.tables.Tables,
		selectScope:       analyzer.scoper.rScope,
		ProjectionErr:     analyzer.projErr,
		Comments:          statement.GetComments(),
	}, nil
}

func (a *analyzer) setError(err error) {
	if len(a.inProjection) > 0 && vterrors.ErrState(err) == vterrors.NonUniqError {
		a.projErr = err
	} else {
		a.err = err
	}
}

// analyzeDown pushes new scopes when we encounter sub queries,
// and resolves the table a column is using
func (a *analyzer) analyzeDown(cursor *sqlparser.Cursor) bool {
	// If we have an error we keep on going down the tree without checking for anything else
	// this way we can abort when we come back up.
	if !a.shouldContinue() {
		return true
	}

	if err := checkForInvalidConstructs(cursor); err != nil {
		a.setError(err)
		return true
	}

	a.scoper.down(cursor)

	n := cursor.Node()
	switch node := n.(type) {
	case sqlparser.SelectExprs:
		if !isParentSelect(cursor) {
			break
		}

		a.inProjection = append(a.inProjection, true)
	case *sqlparser.Order:
		a.analyzeOrderByGroupByExprForLiteral(node.Expr, "order clause")
	case sqlparser.GroupBy:
		for _, grpExpr := range node {
			a.analyzeOrderByGroupByExprForLiteral(grpExpr, "group statement")
		}
	case *sqlparser.ColName:
		tsRecursive, ts, qt, err := a.resolveColumn(node, a.scoper.currentScope())
		if err != nil {
			a.setError(err)
		} else {
			a.binder.exprRecursiveDeps[node] = tsRecursive
			a.binder.exprDeps[node] = ts
			if qt != nil {
				a.exprTypes[node] = *qt
			}
		}
	case *sqlparser.FuncExpr:
		if node.Distinct {
			err := vterrors.Errorf(vtrpcpb.Code_INVALID_ARGUMENT, "syntax error: %s", sqlparser.String(node))
			if len(node.Exprs) != 1 {
				a.setError(err)
			} else if _, ok := node.Exprs[0].(*sqlparser.AliasedExpr); !ok {
				a.setError(err)
			}
		}
	}

	// this is the visitor going down the tree. Returning false here would just not visit the children
	// to the current node, but that is not what we want if we have encountered an error.
	// In order to abort the whole visitation, we have to return true here and then return false in the `analyzeUp` method
	return true
}

func isParentSelect(cursor *sqlparser.Cursor) bool {
	_, isSelect := cursor.Parent().(*sqlparser.Select)
	return isSelect
}

// tableInfoFor returns the table info for the table set. It should contains only single table.
func (a *analyzer) tableInfoFor(id TableSet) (TableInfo, error) {
	numberOfTables := id.NumberOfTables()
	if numberOfTables == 0 {
		return nil, nil
	}
	if numberOfTables > 1 {
		return nil, vterrors.Errorf(vtrpcpb.Code_INTERNAL, "[BUG] should only be used for single tables")
	}
	return a.tables.Tables[id.TableOffset()], nil
}

type originable interface {
	tableSetFor(t *sqlparser.AliasedTableExpr) TableSet
	depsForExpr(expr sqlparser.Expr) (TableSet, *querypb.Type)
}

func (a *analyzer) depsForExpr(expr sqlparser.Expr) (TableSet, *querypb.Type) {
	ts := a.binder.exprRecursiveDeps.Dependencies(expr)
	qt, isFound := a.exprTypes[expr]
	if !isFound {
		return ts, nil
	}
	return ts, &qt
}

func (v *vTableInfo) checkForDuplicates() error {
	for i, name := range v.columnNames {
		for j, name2 := range v.columnNames {
			if i == j {
				continue
			}
			if name == name2 {
				return vterrors.NewErrorf(vtrpcpb.Code_INVALID_ARGUMENT, vterrors.DupFieldName, "Duplicate column name '%s'", name)
			}
		}
	}
	return nil
}

func (a *analyzer) analyze(statement sqlparser.Statement) error {
	_ = sqlparser.Rewrite(statement, a.analyzeDown, a.analyzeUp)
	return a.err
}

func (a *analyzer) analyzeUp(cursor *sqlparser.Cursor) bool {
	if !a.shouldContinue() {
		return false
	}

	if err := a.tables.up(cursor); err != nil {
		a.setError(err)
		return false
	}

	if err := a.scoper.up(cursor); err != nil {
		a.setError(err)
		return false
	}

	switch cursor.Node().(type) {
	case sqlparser.SelectExprs:
		if isParentSelect(cursor) {
			a.popProjection()
		}
	}

	return a.shouldContinue()
}

func checkForInvalidConstructs(cursor *sqlparser.Cursor) error {
	switch node := cursor.Node().(type) {
	case *sqlparser.JoinTableExpr:
		if node.Condition.Using != nil {
			return vterrors.New(vtrpcpb.Code_UNIMPLEMENTED, "unsupported: join with USING(column_list) clause for complex queries")
		}
		if node.Join == sqlparser.NaturalJoinType || node.Join == sqlparser.NaturalRightJoinType || node.Join == sqlparser.NaturalLeftJoinType {
			return vterrors.New(vtrpcpb.Code_UNIMPLEMENTED, "unsupported: "+node.Join.ToString())
		}
	case *sqlparser.Select:
		if node.Having != nil {
			return Gen4NotSupportedF("HAVING")
		}
	case *sqlparser.Subquery:
		return Gen4NotSupportedF("subquery")
	case *sqlparser.FuncExpr:
		if sqlparser.IsLockingFunc(node) {
			return Gen4NotSupportedF("locking functions")
		}
	}

	return nil
}

func createVTableInfoForExpressions(expressions sqlparser.SelectExprs) *vTableInfo {
	vTbl := &vTableInfo{}
	for _, selectExpr := range expressions {
		expr, ok := selectExpr.(*sqlparser.AliasedExpr)
		if !ok {
			continue
		}
		vTbl.cols = append(vTbl.cols, expr.Expr)
		if expr.As.IsEmpty() {
			switch expr := expr.Expr.(type) {
			case *sqlparser.ColName:
				// for projections, we strip out the qualifier and keep only the column name
				vTbl.columnNames = append(vTbl.columnNames, expr.Name.String())
			default:
				vTbl.columnNames = append(vTbl.columnNames, sqlparser.String(expr))
			}
		} else {
			vTbl.columnNames = append(vTbl.columnNames, expr.As.String())
		}
	}
	return vTbl
}

func (a *analyzer) popProjection() {
	a.inProjection = a.inProjection[:len(a.inProjection)-1]
}

func (a *analyzer) shouldContinue() bool {
	return a.err == nil
}

// Gen4NotSupportedF returns a common error for shortcomings in the gen4 planner
func Gen4NotSupportedF(format string, args ...interface{}) error {
	message := fmt.Sprintf("gen4 does not yet support: "+format, args...)

	// add the line that this happens in so it is easy to find it
	stack := string(debug.Stack())
	lines := strings.Split(stack, "\n")
	message += "\n" + lines[6]
	return vterrors.New(vtrpcpb.Code_UNIMPLEMENTED, message)
}
