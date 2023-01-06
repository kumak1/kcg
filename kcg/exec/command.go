package exec

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func IsCommandAvailable(name string) bool {
	if err := exec.Command(name, "--help").Run(); err != nil {
		return false
	}
	return true
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func DirExists(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}

func BranchExists(path string, branch string) bool {
	cmd := exec.Command("git", "show-ref", "-q", "--heads", branch)
	cmd.Dir = path
	err := cmd.Run()
	return err == nil
}

func GhqPath(repo string) (string, error) {
	cmd := exec.Command("ghq", "list", "-p", "-e", repo)
	out, err := cmd.Output()
	path := string(out)
	path = strings.TrimRight(path, "\n")
	return path, err
}

func GhqList() map[string]string {
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
