package ghq

import (
	"github.com/kumak1/kcg/exec"
	"path/filepath"
	"strings"
)

var kcgExec exec.Interface

func initialize() {
	if kcgExec == nil {
		kcgExec = exec.New()
	}
}

func Valid() bool {
	initialize()
	return kcgExec.NotError("", "ghq", "--help")
}

func Get(repo string) (string, error) {
	initialize()
	return kcgExec.Output("", "ghq", "get", repo)
}

func Path(repo string) (string, error) {
	initialize()
	return kcgExec.Output("", "ghq", "list", "-p", "-e", repo)
}

func List() map[string]string {
	initialize()
	ghqList, _ := kcgExec.Output("", "ghq", "list", "-p")
	pathList := map[string]string{}
	for _, path := range strings.Split(ghqList, "\n") {
		if path != "" {
			url, _ := kcgExec.Output(path, "git", "config", "--get", "remote.origin.url")
			organization := filepath.Base(filepath.Dir(path))
			repository := filepath.Base(path)
			pathList[organization+"/"+repository] = url
		}
	}
	return pathList
}
