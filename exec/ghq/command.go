package ghq

import (
	"github.com/kumak1/kcg/exec"
	"strings"
)

var kcgExec exec.Interface

type (
	Interface interface {
		Valid() bool
		Get(repo string) (string, error)
		Path(repo string) (string, error)
		List() []string
	}
	defaultExec struct{}
)

func (d defaultExec) Valid() bool {
	return kcgExec.NotError("", "ghq", "--help")
}

func (d defaultExec) Get(repo string) (string, error) {
	return kcgExec.Output("", "ghq", "get", repo)
}

func (d defaultExec) Path(repo string) (string, error) {
	return kcgExec.Output("", "ghq", "list", "-p", "-e", repo)
}

func (d defaultExec) List() []string {
	list, _ := kcgExec.Output("", "ghq", "list", "-p")
	return strings.Split(list, "\n")
}

func New(e exec.Interface) Interface {
	kcgExec = e
	var i Interface = defaultExec{}
	return i
}
