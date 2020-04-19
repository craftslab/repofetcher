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
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"repofetcher/config"
	"repofetcher/imple"
	"repofetcher/runtime"
)

type Http struct {
	cfg    config.Config
	router *mux.Router
}

func (h *Http) Init(_ *config.Config) error {
	h.router = mux.NewRouter()
	h.router.HandleFunc("/", h.routine).Methods("POST")

	return nil
}

func (h Http) Run(addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: h.router,
	}

	return server.ListenAndServe()
}

func (h Http) routine(resp http.ResponseWriter, req *http.Request) {
	if err := h.read(req); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	buf, err := h.request()
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = runtime.Run(h.operation, buf)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(resp).Encode(map[string]bool{"fetched": true}); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	resp.WriteHeader(http.StatusOK)
}

func (h Http) read(req *http.Request) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &h.cfg); err != nil {
		return err
	}

	return nil
}

func (h Http) request() ([]interface{}, error) {
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

	for _, repo := range h.cfg.Repo {
		for _, clone := range repo.Clone {
			req = append(req, helper(repo, clone))
		}
	}

	if len(req) == 0 {
		return nil, errors.New("request null")
	}

	return req, nil
}

func (h Http) operation(req interface{}) interface{} {
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
