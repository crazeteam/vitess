/*
   Copyright 2014 Outbrain Inc.

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

package inst

import (
	"testing"

	"vitess.io/vitess/go/vt/orchestrator/config"
	test "vitess.io/vitess/go/vt/orchestrator/external/golib/tests"
)

func init() {
	config.Config.HostnameResolveMethod = "none"
	config.MarkConfigurationLoaded()
}

func TestGetAnalysisInstanceType(t *testing.T) {
	{
		analysis := &ReplicationAnalysis{}
		test.S(t).ExpectEquals(string(analysis.GetAnalysisInstanceType()), "intermediate-primary")
	}
	{
		analysis := &ReplicationAnalysis{IsPrimary: true}
		test.S(t).ExpectEquals(string(analysis.GetAnalysisInstanceType()), "primary")
	}
	{
		analysis := &ReplicationAnalysis{IsCoPrimary: true}
		test.S(t).ExpectEquals(string(analysis.GetAnalysisInstanceType()), "co-primary")
	}
	{
		analysis := &ReplicationAnalysis{IsPrimary: true, IsCoPrimary: true}
		test.S(t).ExpectEquals(string(analysis.GetAnalysisInstanceType()), "co-primary")
	}
}
