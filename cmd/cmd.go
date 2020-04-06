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

package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"

	"repofetcher/config"
	"repofetcher/fetcher"
)

const (
	totalAddr = 2
)

var (
	app     = kingpin.New("repofetcher", "Repository Fetcher").Author(Author).Version(Version)
	addr    = app.Flag("addr", "Server listen address (http)").Default(":9093").String()
	mode    = app.Flag("mode", "Communication mode (http|stdio)").Default("stdio").String()
	repo    = app.Flag("repo", "Repo list in json (stdio)").Default("repo.json").String()
	routine = app.Flag("routine", "Routine to fulfill requests (http)").Default("0").Int()
)

func Run() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	addr, err := parseAddr(*addr)
	if err != nil {
		log.Fatal("addr invalid: ", err.Error())
	}

	mode, err := parseMode(*mode)
	if err != nil {
		log.Fatal("mode invalid: ", err.Error())
	}

	cfg, err := parseRepo(*repo)
	if err != nil {
		log.Fatal("repo invalid: ", err.Error())
	}

	routine, err := parseRoutine(*routine)
	if err != nil {
		log.Fatal("routine invalid: ", err.Error())
	}

	if err := runFetcher(addr, mode, &cfg, routine); err != nil {
		log.Fatal("fetcher failed: ", err.Error())
	}

	log.Println("fetcher completed.")
}

func parseAddr(data string) (string, error) {
	buf := strings.Split(data, ":")
	if len(buf) != totalAddr {
		return "", errors.New("separator invalid")
	}

	port := strings.TrimSpace(buf[1])
	if port == "" {
		return "", errors.New("port invalid")
	}

	p, err := strconv.Atoi(port)
	if err != nil || p <= 0 {
		return "", errors.Wrap(err, "port invalid")
	}

	return data, nil
}

func parseMode(data string) (string, error) {
	matched := false
	for _, val := range fetcher.Mode {
		if val == data {
			matched = true
			break
		}
	}

	if !matched {
		return "", errors.New("mode mismatched")
	}

	return data, nil
}

func parseRepo(name string) (config.Config, error) {
	cfg := config.Config{}

	file, err := os.Open(name)
	if err != nil {
		return cfg, errors.Wrap(err, "open failed")
	}

	buf, _ := ioutil.ReadAll(file)

	if err := file.Close(); err != nil {
		return cfg, errors.Wrap(err, "close failed")
	}

	if err := json.Unmarshal(buf, &cfg); err != nil {
		return cfg, errors.Wrap(err, "unmarshal failed")
	}

	return cfg, nil
}

func parseRoutine(data int) (int, error) {
	if data <= 0 {
		data = runtime.NumCPU() / 2
		if data <= 0 {
			data = 1
		}
	}

	return data, nil
}

func runFetcher(addr, mode string, cfg *config.Config, routine int) error {
	return fetcher.Run(addr, mode, cfg, routine)
}
