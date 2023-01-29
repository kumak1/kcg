package git

import (
	"os/exec"
	"strings"
)

func BranchExists(path string, branch string) bool {
	cmd := exec.Command("git", "show-ref", "-q", "--heads", branch)
	cmd.Dir = path
	return cmd.Run() == nil
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

func Switch(path string, branch string) (string, error) {
	cmd := exec.Command("git", "switch", branch)
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	return strings.TrimRight(string(out), "\n"), err
}

func Pull(path string) (string, error) {
	cmd := exec.Command("git", "pull")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	return strings.TrimRight(string(out), "\n"), err
}

func Clone(repo string, path string) (string, error) {
	cmd := exec.Command("git", "clone", repo, path)
	out, err := cmd.CombinedOutput()
	return strings.TrimRight(string(out), "\n"), err
}

func Cleanup(path string) (string, error) {
	cmd := exec.Command("sh", "-c", "git branch --merged|egrep -v '\\*|develop|main|master'|xargs git branch -d")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	return strings.TrimRight(string(out), "\n"), err
}
