package git

import (
	"github.com/kumak1/kcg/exec"
)

var kcgExec exec.Interface

func initialize() {
	if kcgExec == nil {
		kcgExec = exec.New()
	}
}

func BranchExists(path string, branch string) bool {
	initialize()
	return kcgExec.NotError(path, "git", "show-ref", "-q", "--heads", branch)
}

func CurrentBranchName(path string) string {
	initialize()
	if out, err := kcgExec.Output(path, "git", "rev-parse", "--abbrev-ref", "HEAD"); err == nil {
		return out
	} else {
		return ""
	}
}

func Switch(path string, branch string) (string, error) {
	initialize()
	return kcgExec.Output(path, "git", "switch", branch)
}

func Pull(path string) (string, error) {
	initialize()
	return kcgExec.Output(path, "git", "pull")
}

func Clone(repo string, path string) (string, error) {
	initialize()
	return kcgExec.Output(path, "git", "clone", repo, path)
}

func Cleanup(path string) (string, error) {
	initialize()
	return kcgExec.Output(path, "sh", "-c", "git branch --merged|egrep -v '\\*|develop|main|master'|xargs git branch -d")
}
