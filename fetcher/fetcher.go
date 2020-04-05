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
	"github.com/pkg/errors"
)

type Fetcher interface {
	Init(config map[string]interface{}) error
	Run(addr string, repo map[string]interface{}, routine int) error
}

var (
	Mode = initMode()
)

var (
	fetchers = map[string]Fetcher{
		"http":  &Http{},
		"stdio": &StdIo{},
	}
)

func Run(addr, mode string, repo map[string]interface{}, routine int) error {
	return runFetcher(fetchers[mode], addr, repo, routine)
}

func initMode() []string {
	var buf []string

	for key := range fetchers {
		buf = append(buf, key)
	}

	return buf
}

func runFetcher(f Fetcher, addr string, repo map[string]interface{}, routine int) error {
	if err := f.Init(nil); err != nil {
		return errors.Wrap(err, "init failed")
	}

	return f.Run(addr, repo, routine)
}
