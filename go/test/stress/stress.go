/*
Copyright 2021 The Vitess Authors.

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

package stress

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"vitess.io/vitess/go/mysql"
)

const (
	templateNewTable = `create table %s (
	id bigint,
	val varchar(64),
	primary key (id)
) Engine=InnoDB
`
)

type (
	Result struct {
		countSelect       int
		countSelectFailed int
		countInsert       int
		countInsertFailed int
	}

	table struct {
		name         string
		rows, nextID int
		mu           sync.Mutex
	}

	Stresser struct {
		doneCh     chan Result
		tbls       []*table
		connParams *mysql.ConnParams
		maxClient  int
		duration   time.Duration
		t          *testing.T
		finish     bool
	}
)

func (r Result) PrintQPS(seconds float64) {
	fmt.Printf(`QPS:
	select: %d (failed: %d)
	insert: %d (failed: %d)
	---------
	total:	%d (failed: %d)
	
Queries:
	select: %d (failed: %d)
	insert: %d (failed: %d)
	---------
	total:	%d (failed: %d)
	
`, r.countSelect/int(seconds), r.countSelectFailed/int(seconds),
		r.countInsert/int(seconds), r.countInsertFailed/int(seconds),
		(r.countInsert+r.countSelect)/int(seconds), (r.countInsertFailed+r.countSelectFailed)/int(seconds),
		r.countSelect, r.countSelectFailed,
		r.countInsert, r.countInsertFailed,
		r.countInsert+r.countSelect, r.countInsertFailed+r.countSelectFailed)
}

func generateNewTables(nb int) []*table {
	tbls := make([]*table, 0, nb)
	for i := 0; i < nb; i++ {
		tbls = append(tbls, &table{
			name: fmt.Sprintf("stress_t%d", i),
		})
	}
	return tbls
}

func createTables(t *testing.T, params *mysql.ConnParams, nb int) []*table {
	conn := newClient(t, params)
	defer conn.Close()

	tbls := generateNewTables(nb)
	for _, tbl := range tbls {
		exec(t, conn, fmt.Sprintf(templateNewTable, tbl.name))
	}
	return tbls
}

func (s *Stresser) SetConn(conn *mysql.ConnParams) *Stresser {
	s.connParams = conn
	return s
}

func New(t *testing.T, conn *mysql.ConnParams, duration time.Duration) *Stresser {
	return &Stresser{
		doneCh:     make(chan Result),
		t:          t,
		connParams: conn,
		duration:   duration,
		maxClient:  10,
	}
}

func (s *Stresser) Start() *Stresser {
	fmt.Println("Starting load testing ...")
	s.tbls = createTables(s.t, s.connParams, 100)
	go s.startClients()
	return s
}

func (s *Stresser) startClients() {
	resultCh := make(chan Result, s.maxClient)
	for i := 0; i < s.maxClient; i++ {
		go s.startStressClient(resultCh)
	}

	perClientResults := make([]Result, 0, s.maxClient)
	for i := 0; i < s.maxClient; i++ {
		newResult := <-resultCh
		perClientResults = append(perClientResults, newResult)
	}

	var finalResult Result
	for _, r := range perClientResults {
		finalResult.countSelect += r.countSelect
		finalResult.countSelectFailed += r.countSelectFailed
		finalResult.countInsert += r.countInsert
		finalResult.countInsertFailed += r.countInsertFailed
	}
	s.doneCh <- finalResult
}

func (s *Stresser) startStressClient(resultCh chan Result) {
	connParams := s.connParams
	conn := newClient(s.t, connParams)
	defer conn.Close()

	var res Result

	timeout := time.After(s.duration)

outer:
	for !s.finish {
		if connParams != s.connParams {
			connParams = s.connParams
			conn.Close()
			conn = newClient(s.t, connParams)
		}
		select {
		case <-timeout:
			break outer
		case <-time.After(15 * time.Microsecond):
			s.insertToRandomTable(conn, &res)
		case <-time.After(1 * time.Microsecond):
			s.selectFromRandomTable(conn, &res)
		}
	}
	resultCh <- res
}

func (s *Stresser) insertToRandomTable(conn *mysql.Conn, r *Result) {
	tblI := rand.Int() % len(s.tbls)
	s.tbls[tblI].mu.Lock()
	defer s.tbls[tblI].mu.Unlock()

	query := fmt.Sprintf("insert into %s(id, val) values(%d, 'name')", s.tbls[tblI].name, s.tbls[tblI].nextID)
	if exec(s.t, conn, query) != nil {
		s.tbls[tblI].nextID++
		s.tbls[tblI].rows++
		r.countInsert++
	} else {
		r.countInsertFailed++
	}
}

func (s *Stresser) selectFromRandomTable(conn *mysql.Conn, r *Result) {
	tblI := rand.Int() % len(s.tbls)
	s.tbls[tblI].mu.Lock()
	defer s.tbls[tblI].mu.Unlock()

	query := fmt.Sprintf("select * from %s limit 500", s.tbls[tblI].name)
	expLength := s.tbls[tblI].rows
	if expLength > 500 {
		expLength = 500
	}
	if assertLength(s.t, conn, query, expLength) {
		r.countSelect++
	} else {
		r.countSelectFailed++
	}
}

func (s *Stresser) Wait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	select {
	case res := <-s.doneCh:
		res.PrintQPS(s.duration.Seconds())
	case <-timeoutCh:
		s.finish = true
		res := <-s.doneCh
		res.PrintQPS(s.duration.Seconds())
	}
}
