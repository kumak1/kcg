package internal

import (
	"os"
	"os/exec"
	"strings"
)

type (
	Interface interface {
		FileExists(string) bool
		DirExists(string) bool
		Output(string, string, ...string) (string, error)
		NotError(string, string, ...string) bool
	}
	defaultExec struct{}
)

func New() Interface {
	var i Interface = defaultExec{}
	return i
}

func (e defaultExec) FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func (e defaultExec) DirExists(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}

func (e defaultExec) Output(path string, name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	return strings.TrimRight(string(out), "\n"), err
}

func (e defaultExec) NotError(path string, name string, arg ...string) bool {
	cmd := exec.Command(name, arg...)
	cmd.Dir = path
	return cmd.Run() == nil
}
