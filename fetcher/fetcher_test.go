// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fetcher

import (
	"testing"
)

func TestInitMode(t *testing.T) {
	if buf := initMode(); len(buf) == 0 {
		t.Error("FAIL")
	}
}

func TestRunFetcher(t *testing.T) {
	repo := map[string]interface{}{
		"repo": []map[string]interface{}{
			{
				"branch": "master",
				"clone": []map[string]interface{}{
					{
						"label": "default",
						"sparse": []string{
							"cmd",
						},
					},
				},
				"depth": 1,
				"name":  "repofetcher",
				"path":  "repofetcher",
				"url":   "https://github.com/craftslab/repofetcher.git",
			},
		},
	}

	if err := runFetcher(&StdIo{}, "", repo, 0); err != nil {
		t.Error("FAIL:", err)
	}
}
