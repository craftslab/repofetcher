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

package config

type Config struct {
	Repo []Repo `json:"repo"`
}

type Repo struct {
	Branch string  `json:"branch"`
	Clone  []Clone `json:"clone"`
	Depth  int     `json:"depth"`
	Name   string  `json:"name"`
	Path   string  `json:"path"`
	Url    string  `json:"url"`
}

type Clone struct {
	Sparse []string `json:"sparse"`
}
