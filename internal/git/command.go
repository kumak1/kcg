package git

import "github.com/kumak1/kcg/internal"

type (
	Interface interface {
		BranchExists(path string, branch string) bool
		CurrentBranchName(path string) string
		Switch(path string, branch string) (string, error)
		Pull(path string) (string, error)
		Clone(repo string, path string) (string, error)
		Cleanup(path string) (string, error)
		OriginUrl(path string) (string, error)
	}
	defaultExec struct{}
)

var kcgExec internal.Interface

func (d defaultExec) BranchExists(path string, branch string) bool {
	return kcgExec.NotError(path, "git", "show-ref", "-q", "--heads", branch)
}

func (d defaultExec) CurrentBranchName(path string) string {
	if out, err := kcgExec.Output(path, "git", "rev-parse", "--abbrev-ref", "HEAD"); err == nil {
		return out
	} else {
		return ""
	}
}

func (d defaultExec) Switch(path string, branch string) (string, error) {
	return kcgExec.Output(path, "git", "switch", branch)
}

func (d defaultExec) Pull(path string) (string, error) {
	return kcgExec.Output(path, "git", "pull")
}

func (d defaultExec) Clone(repo string, path string) (string, error) {
	return kcgExec.Output(path, "git", "clone", repo, path)
}

func (d defaultExec) Cleanup(path string) (string, error) {
	return kcgExec.Output(path, "sh", "-c", "git branch --merged|egrep -v '\\*|develop|main|master'|xargs git branch -d")
}

func (d defaultExec) OriginUrl(path string) (string, error) {
	return kcgExec.Output(path, "git", "config", "--get", "remote.origin.url")
}

func New(e internal.Interface) Interface {
	kcgExec = e
	var i Interface = defaultExec{}
	return i
}
