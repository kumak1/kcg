package ghq

import (
	"os/exec"
	"path/filepath"
	"strings"
)

func Valid() bool {
	return exec.Command("ghq", "--help").Run() == nil
}

func Get(repo string) (string, error) {
	out, err := exec.Command("ghq", "get", repo).CombinedOutput()
	return strings.TrimRight(string(out), "\n"), err
}

func Path(repo string) (string, error) {
	out, err := exec.Command("ghq", "list", "-p", "-e", repo).Output()
	return strings.TrimRight(string(out), "\n"), err
}

func List() map[string]string {
	ghqList, _ := exec.Command("ghq", "list", "-p").Output()
	pathList := map[string]string{}
	for _, path := range strings.Split(string(ghqList), "\n") {
		if path != "" {
			cmd := exec.Command("git", "config", "--get", "remote.origin.url")
			cmd.Dir = path
			url, _ := cmd.Output()
			organization := filepath.Base(filepath.Dir(path))
			repository := filepath.Base(path)
			pathList[organization+"/"+repository] = strings.TrimRight(string(url), "\n")
		}
	}
	return pathList
}
