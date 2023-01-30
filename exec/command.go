package exec

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
	var i Interface
	i = defaultExec{}
	return i
}

func (e defaultExec) FileExists(path string) bool {
	return fileExists(path)
}

func (e defaultExec) DirExists(path string) bool {
	return dirExists(path)
}

func (e defaultExec) Output(path string, name string, arg ...string) (string, error) {
	return output(path, name, arg...)
}

func (e defaultExec) NotError(path string, name string, arg ...string) bool {
	return notError(path, name, arg...)
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func dirExists(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}

func output(path string, name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	return strings.TrimRight(string(out), "\n"), err
}

func notError(path string, name string, arg ...string) bool {
	cmd := exec.Command(name, arg...)
	cmd.Dir = path
	return cmd.Run() == nil
}
