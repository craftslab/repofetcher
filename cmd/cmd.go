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
	app  = kingpin.New("repofetcher", "Repository Fetcher").Author(Author).Version(Version)
	addr = app.Flag("addr", "Server listen address (http)").Default(":9093").String()
	mode = app.Flag("mode", "Communication mode (http|stdio)").Default("stdio").String()
	repo = app.Flag("repo", "Repo list in json (stdio)").Default("repo.json").String()
)

func Run() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	mode, err := parseMode(*mode)
	if err != nil {
		log.Fatal("mode invalid: ", err.Error())
	}

	var _addr string
	var cfg config.Config

	switch mode {
	case "http":
		_addr, err = parseAddr(*addr)
	case "stdio":
		cfg, err = parseConfig(*repo)
	default:
		err = errors.New("mode invalid")
	}

	if err != nil {
		log.Fatal("addr or config failed: ", err.Error())
	}

	if err := runFetcher(mode, _addr, &cfg); err != nil {
		log.Fatal("fetcher failed: ", err.Error())
	}

	log.Println("fetcher completed.")
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

func parseConfig(name string) (config.Config, error) {
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

func runFetcher(mode, addr string, cfg *config.Config) error {
	return fetcher.Run(mode, addr, cfg)
}
