package exec

import (
	"os"
	"os/exec"
	"strings"
)

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

func Output(path string, name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	return strings.TrimRight(string(out), "\n"), err
}

func NotError(path string, arg ...string) bool {
	return exec.Command("ghq", "--help").Run() == nil
}
