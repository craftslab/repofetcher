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
	"repofetcher/imple"
	"repofetcher/runtime"
)

type StdIo struct {
	cfg config.Config
}

func (s *StdIo) Init(cfg *config.Config) error {
	s.cfg = *cfg
	return nil
}

func (s StdIo) Run(_ string) error {
	return s.routine()
}

func (s StdIo) routine() error {
	req, err := s.request()
	if err != nil {
		return errors.Wrap(err, "request invalid")
	}

	_, err = runtime.Run(s.operation, req)

	return err
}

func (s StdIo) request() ([]interface{}, error) {
	helper := func(repo config.Repo, clone config.Clone) config.Repo {
		return config.Repo{
			Branch: repo.Branch,
			Clone:  []config.Clone{clone},
			Depth:  repo.Depth,
			Name:   repo.Name,
			Path:   repo.Path,
			Url:    repo.Url,
		}
	}

	var req []interface{}

	for _, repo := range s.cfg.Repo {
		for _, clone := range repo.Clone {
			req = append(req, helper(repo, clone))
		}
	}

	if len(req) == 0 {
		return nil, errors.New("request null")
	}

	return req, nil
}

func (s StdIo) operation(req interface{}) interface{} {
	i := imple.Imple{}

	if err := i.Check(); err != nil {
		return nil
	}

	repo := req.(config.Repo)
	clone := repo.Clone[0]
	sparse := false

	if len(clone.Sparse) != 0 {
		sparse = true
	}

	if err := i.Init(repo.Branch, repo.Name, repo.Path, repo.Url, repo.Depth, sparse); err != nil {
		return nil
	}

	if err := i.Clone(); err != nil {
		return nil
	}

	if sparse {
		for _, item := range clone.Sparse {
			_ = i.Add(item)
		}
	}

	return nil
}
