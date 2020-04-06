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

package imple

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	majorVersion = 2
	minorVersion = 26
)

type Imple struct {
	branch string
	depth  int
	name   string
	path   string
	sparse bool
	url    string
}

func (i *Imple) Init(branch, name, path, url string, depth int, sparse bool) error {
	if branch == "" || name == "" || path == "" || url == "" || depth <= 0 {
		return errors.New("config invalid")
	}

	i.branch = branch
	i.depth = depth
	i.name = name
	i.path = path
	i.sparse = sparse
	i.url = url

	return nil
}

func (i Imple) Add(data string) error {
	if !i.sparse {
		return errors.New("sparse required")
	}

	dir, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "getwd failed")
	}

	if err := os.Chdir(i.path); err != nil {
		return errors.Wrap(err, "chdir failed")
	}

	cmd := exec.Command("git", "sparse-checkout", "init", "--cone")
	_, err = cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "init failed")
	}

	cmd = exec.Command("git", "sparse-checkout", "add", data)
	_, err = cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "add failed")
	}

	if err := os.Chdir(dir); err != nil {
		return errors.Wrap(err, "chdir failed")
	}

	return nil
}

func (i Imple) Check() error {
	_, err := exec.LookPath("git")
	if err != nil {
		return errors.Wrap(err, "command not found")
	}

	cmd := exec.Command("git", "version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "version failed")
	}

	buf := strings.Trim(string(out), "git version")
	version := strings.Split(buf, ".")

	if len(version) < 3 {
		return errors.New("version invalid")
	}

	major, err := strconv.Atoi(version[0])
	if err != nil {
		return errors.Wrap(err, "major version invalid")
	}

	minor, err := strconv.Atoi(version[1])
	if err != nil {
		return errors.Wrap(err, "minor version invalid")
	}

	if major < majorVersion || minor < minorVersion {
		return errors.New(strconv.Itoa(majorVersion) + "." + strconv.Itoa(minorVersion) + "+ required")
	}

	return nil
}

func (i Imple) Clone() error {
	var cmd *exec.Cmd

	if i.sparse {
		cmd = exec.Command("git", "clone",
			"-b", i.branch,
			"--depth", strconv.Itoa(i.depth),
			"--filter=blob:none",
			"--quiet",
			"--recurse-submodules",
			"--sparse",
			i.url+"/"+i.name,
			i.path)
	} else {
		cmd = exec.Command("git", "clone",
			"-b", i.branch,
			"--depth", strconv.Itoa(i.depth),
			"--quiet",
			"--recurse-submodules",
			i.url+"/"+i.name,
			i.path)
	}

	_, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "clone failed")
	}

	return nil
}
