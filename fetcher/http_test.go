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
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"repofetcher/config"
)

func TestRead(t *testing.T) {
	cfg := config.Config{
		Repo: make([]config.Repo, 1),
	}

	cfg.Repo[0] = config.Repo{
		Branch: "master",
		Clone:  make([]config.Clone, 1),
		Depth:  1,
		Name:   "repofetcher",
		Path:   "repofetcher",
		Url:    "https://github.com/craftslab",
	}

	cfg.Repo[0].Clone[0] = config.Clone{
		Sparse: []string{
			"cmd",
		},
	}

	body, _ := json.Marshal(cfg)
	req := httptest.NewRequest("POST", "http://localhost:9093", bytes.NewBuffer(body))

	h := &Http{}
	if err := h.read(req); err != nil {
		t.Error("FAIL")
	}
}
