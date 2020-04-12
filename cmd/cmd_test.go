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
	"testing"
)

func TestParseAddr(t *testing.T) {
	if _, err := parseAddr(""); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseAddr("127.0.0.1"); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseAddr(":"); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseAddr("127.0.0.1:"); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseAddr(":9093"); err != nil {
		t.Error("FAIL")
	}
	if _, err := parseAddr("127.0.0.1:9093"); err != nil {
		t.Error("FAIL")
	}
}

func TestParseMode(t *testing.T) {
	if _, err := parseMode("invalid"); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseMode("http"); err != nil {
		t.Error("FAIL")
	}

	if _, err := parseMode("stdio"); err != nil {
		t.Error("FAIL")
	}
}

func TestParseRepo(t *testing.T) {
	if _, err := parseRepo(""); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseRepo("../test/repo.json"); err != nil {
		t.Error("FAIL")
	}
}
