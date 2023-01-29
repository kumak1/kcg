package ghq

import (
	"github.com/kumak1/kcg/exec"
	"path/filepath"
	"strings"
)

func Valid() bool {
	return exec.NotError("", "ghq", "--help")
}

func Get(repo string) (string, error) {
	return exec.Output("", "ghq", "get", repo)
}

func Path(repo string) (string, error) {
	return exec.Output("", "ghq", "list", "-p", "-e", repo)
}

func List() map[string]string {
	ghqList, _ := exec.Output("", "ghq", "list", "-p")
	pathList := map[string]string{}
	for _, path := range strings.Split(ghqList, "\n") {
		if path != "" {
			url, _ := exec.Output(path, "git", "config", "--get", "remote.origin.url")
			organization := filepath.Base(filepath.Dir(path))
			repository := filepath.Base(path)
			pathList[organization+"/"+repository] = url
		}
	}
	return pathList
}
