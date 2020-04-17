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

	"repofetcher/config"
)

type Fetcher interface {
	Init(cfg *config.Config) error
	Run(addr string) error
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

func Run(addr, mode string, cfg *config.Config) error {
	return runFetcher(fetchers[mode], addr, cfg)
}

func initMode() []string {
	var buf []string

	for key := range fetchers {
		buf = append(buf, key)
	}

	return buf
}

func runFetcher(f Fetcher, addr string, cfg *config.Config) error {
	if err := f.Init(cfg); err != nil {
		return errors.Wrap(err, "init failed")
	}

	return f.Run(addr)
}
