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
	"repofetcher/config"
)

type Http struct {
	cfg config.Config
}

func (h *Http) Init(cfg *config.Config) error {
	h.cfg = *cfg
	return nil
}

func (h Http) Run(addr string) error {
	return h.runHttp(addr)
}

func (h Http) runHttp(addr string) error {
	// TODO
	return nil
}
