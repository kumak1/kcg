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

func BranchExists(path string, branch string) bool {
	cmd := exec.Command("git", "show-ref", "-q", "--heads", branch)
	cmd.Dir = path
	err := cmd.Run()
	return err == nil
}

func CurrentBranchName(path string) string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = path
	if out, err := cmd.Output(); err == nil {
		return strings.TrimRight(string(out), "\n")
	} else {
		return ""
	}
}
