package git

import (
	"github.com/kumak1/kcg/exec"
)

func BranchExists(path string, branch string) bool {
	return exec.NotError(path, "git", "show-ref", "-q", "--heads", branch)
}

func CurrentBranchName(path string) string {
	if out, err := exec.Output(path, "git", "rev-parse", "--abbrev-ref", "HEAD"); err == nil {
		return out
	} else {
		return ""
	}
}

func Switch(path string, branch string) (string, error) {
	return exec.Output(path, "git", "switch", branch)
}

func Pull(path string) (string, error) {
	return exec.Output(path, "git", "pull")
}

func Clone(repo string, path string) (string, error) {
	return exec.Output(path, "git", "clone", repo, path)
}

func Cleanup(path string) (string, error) {
	return exec.Output(path, "sh", "-c", "git branch --merged|egrep -v '\\*|develop|main|master'|xargs git branch -d")
}
