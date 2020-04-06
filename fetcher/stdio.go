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
	"repofetcher/runtime"
)

type Request struct {
	repo config.Repo
}

type StdIo struct {
	cfg config.Config
}

func (s *StdIo) Init(cfg *config.Config) error {
	s.cfg = *cfg
	return nil
}

func (s StdIo) Run(_ string, _ int) error {
	return s.runStdIo()
}

func (s StdIo) runStdIo() error {
	req, err := s.request()
	if err != nil {
		return errors.Wrap(err, "request invalid")
	}

	_, err = runtime.Run(s.operation, req)

	return err
}

func (s StdIo) request() ([]interface{}, error) {
	helper := func(data config.Repo) ([]Request, error) {
		var req []Request
		for _, item := range data.Clone {
			r := Request{
				repo: config.Repo{
					Branch: data.Branch,
					Clone:  []config.Clone{item},
					Depth:  data.Depth,
					Name:   data.Name,
					Path:   data.Path,
					Url:    data.Url,
				},
			}
			req = append(req, r)
		}
		return req, nil
	}

	var req []interface{}

	for _, item := range s.cfg.Repo {
		r, err := helper(item)
		if err != nil {
			return nil, err
		}
		req = append(req, r)
	}

	if len(req) == 0 {
		return nil, errors.New("request null")
	}

	return req, nil
}

func (s StdIo) operation(req interface{}) interface{} {
	return nil
}
