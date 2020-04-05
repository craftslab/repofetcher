# repofetcher

[![Build Status](https://travis-ci.com/craftslab/repofetcher.svg?branch=master)](https://travis-ci.com/craftslab/repofetcher)
[![Coverage Status](https://coveralls.io/repos/github/craftslab/repofetcher/badge.svg?branch=master)](https://coveralls.io/github/craftslab/repofetcher?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/craftslab/repofetcher)](https://goreportcard.com/report/github.com/craftslab/repofetcher)
[![License](https://img.shields.io/github/license/craftslab/repofetcher.svg?color=brightgreen)](https://github.com/craftslab/repofetcher/blob/master/LICENSE)
[![Tag](https://img.shields.io/github/tag/craftslab/repofetcher.svg?color=brightgreen)](https://github.com/craftslab/repofetcher/tags)



## Introduction

*Repo Fetcher* is a repository fetcher written in Go.



## Feature

- Support to fetch repository in server mode (http).

- Support to fetch repository in standalone mode (stdio).



## Requirement

- Git 2.26+



## Usage

```bash
usage: repofetcher [<flags>]

Repository Fetcher

Flags:
  --help              Show context-sensitive help (also try --help-long and --help-man).
  --version           Show application version.
  --addr=":9093"      Server listen address (http)
  --mode="stdio"      Communication mode (http|stdio)
  --repo="repo.json"  Repo list in json (stdio)
  --routine=0         Routine to fulfill requests (http)
```



## Setting

*Repo Fetcher* parameters can be set in the directory config or as HTTP request.

An example of configuration in [repo.json](https://github.com/craftslab/repofetcher/blob/master/config/repo.json):

```json
{
  "repo": [
    {
      "branch": "master",
      "clone": [
        {
          "label": "default",
          "sparse": [
            "cmd",
            "config"
          ]
        }
      ],
      "depth": 1,
      "name": "repofetcher",
      "path": "repofetcher",
      "url": "https://github.com/craftslab/repofetcher.git"
    }
  ]
}
```



## Running

- Server mode (http)

```bash
repofetcher --mode="http" --addr=":9093" --routine=1
```

- Standalone mode (stdio)

```bash
repofetcher --mode="stdio" --repo="repo.json"
```



## License

Project License can be found [here](LICENSE).
