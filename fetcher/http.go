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
	"net/http"

	"github.com/gorilla/mux"

	"repofetcher/config"
)

type Http struct {
	router *mux.Router
}

func (h *Http) Init(_ *config.Config) error {
	return nil
}

func (h Http) Run(addr string) error {
	h.router = mux.NewRouter()
	h.router.HandleFunc("/", h.runHttp).Methods("POST")

	srv := &http.Server{
		Handler: h.router,
		Addr:    addr,
	}

	return srv.ListenAndServe()
}

func (h Http) runHttp(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
}
