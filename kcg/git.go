package kcg

import (
	"fmt"
	kcgExec "github.com/kumak1/kcg/kcg/exec"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	repositoryConfig map[string]*RepositoryConfig
	standardOut      = os.Stdout
	standardError    = os.Stderr
)

func Command(config Config) IGitOperator {
	repositoryConfig = config.Repos

	var operator IGitOperator
	if config.Ghq {
		operator = ghq{}
	} else {
		operator = git{}
	}

	return operator
}

type IGitOperator interface {
	Cleanup(*RepositoryConfig) ([]byte, error)
	Clone(*RepositoryConfig) ([]byte, error)
	CurrentBranch(*RepositoryConfig) string
	List(string, string) map[string]*RepositoryConfig
	Path(*RepositoryConfig) (string, bool)
	Pull(*RepositoryConfig) ([]byte, error)
	Run(*RepositoryConfig, string) (string, error)
	Switch(*RepositoryConfig, string) ([]byte, error)
}

type git struct{}

type ghq struct{}

func (g git) Cleanup(config *RepositoryConfig) ([]byte, error) {
	if path, exists := g.Path(config); exists {
		return cleanup(path)
	} else {
		return nil, fmt.Errorf("\x1b[31m%s\x1b[0m %s", "invalid path", path)
	}
}

func (g ghq) Cleanup(config *RepositoryConfig) ([]byte, error) {
	if path, exists := g.Path(config); exists {
		return cleanup(path)
	} else {
		return nil, fmt.Errorf("\x1b[31m%s\x1b[0m %s", "invalid path", path)
	}
}

func cleanup(path string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", "git branch --merged|egrep -v '\\*|develop|main|master'|xargs git branch -d")
	cmd.Dir = path
	return cmd.CombinedOutput()
}

func (g git) Clone(config *RepositoryConfig) ([]byte, error) {
	if config.Repo == "" {
		return nil, fmt.Errorf("\x1b[31m%s\x1b[0m %s", "error", "repo is empty")
	}

	if path, exists := g.Path(config); exists {
		return nil, fmt.Errorf("\x1b[33m%s\x1b[0m %s", "exists", path)
	} else {
		cmd := exec.Command("git", "clone", config.Repo, path)
		return cmd.CombinedOutput()
	}
}

func (g ghq) Clone(config *RepositoryConfig) ([]byte, error) {
	if config.Repo == "" {
		return nil, fmt.Errorf("\x1b[31m%s\x1b[0m %s", "error", "repo is empty")
	}

	cmd := exec.Command("ghq", "get", config.Repo)
	return cmd.CombinedOutput()
}

func (g git) CurrentBranch(config *RepositoryConfig) string {
	if path, exists := g.Path(config); exists {
		return kcgExec.CurrentBranchName(path)
	} else {
		return ""
	}
}

func (g ghq) CurrentBranch(config *RepositoryConfig) string {
	if path, exists := g.Path(config); exists {
		return kcgExec.CurrentBranchName(path)
	} else {
		return ""
	}
}

func (g git) List(group string, filter string) map[string]*RepositoryConfig {
	return list(group, filter)
}

func (g ghq) List(group string, filter string) map[string]*RepositoryConfig {
	return list(group, filter)
}

func list(group string, filter string) map[string]*RepositoryConfig {
	repositoryConfigs := map[string]*RepositoryConfig{}
	for index, repo := range repositoryConfig {
		if validGroup(group, repo.Group) && validFilter(filter, index) {
			repositoryConfigs[index] = repo
		}
	}
	return repositoryConfigs
}

func (g git) Path(config *RepositoryConfig) (string, bool) {
	return config.Path, config.Path != "" && kcgExec.FileExists(config.Path)
}

func (g ghq) Path(config *RepositoryConfig) (string, bool) {
	if config.Path != "" {
		return config.Path, kcgExec.FileExists(config.Path)
	} else {
		if config.Repo == "" {
			return "", false
		}
		path, _ := kcgExec.GhqPath(config.Repo)
		return path, path != "" && kcgExec.FileExists(path)
	}
}

func (g git) Pull(config *RepositoryConfig) ([]byte, error) {
	if path, exists := g.Path(config); exists {
		return pull(path)
	} else {
		return nil, fmt.Errorf("\x1b[31m%s\x1b[0m %s", "invalid path", path)
	}
}

func (g ghq) Pull(config *RepositoryConfig) ([]byte, error) {
	if path, exists := g.Path(config); exists {
		return pull(path)
	} else {
		return nil, fmt.Errorf("\x1b[31m%s\x1b[0m %s", "invalid path", path)
	}
}

func pull(path string) ([]byte, error) {
	cmd := exec.Command("git", "pull")
	cmd.Dir = path
	return cmd.CombinedOutput()
}

func (g git) Run(config *RepositoryConfig, command string) (string, error) {
	if path, exists := g.Path(config); exists {
		return run(path, command)
	} else {
		return "", fmt.Errorf("\x1b[31m%s\x1b[0m %s", "invalid path", path)
	}
}

func (g ghq) Run(config *RepositoryConfig, command string) (string, error) {
	if path, exists := g.Path(config); exists {
		return run(path, command)
	} else {
		return "", fmt.Errorf("\x1b[31m%s\x1b[0m %s", "invalid path", path)
	}
}

func run(path string, command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	return strings.TrimRight(string(out), "\n"), err
}

func (g git) Switch(config *RepositoryConfig, branch string) ([]byte, error) {
	if path, exists := g.Path(config); exists {
		return switchBranch(path, convertedBranch(config.Alias, branch))
	} else {
		return nil, fmt.Errorf("\x1b[31m%s\x1b[0m %s", "invalid path", path)
	}
}

func (g ghq) Switch(config *RepositoryConfig, branch string) ([]byte, error) {
	if path, exists := g.Path(config); exists {
		return switchBranch(path, convertedBranch(config.Alias, branch))
	} else {
		return nil, fmt.Errorf("\x1b[31m%s\x1b[0m %s", "invalid path", path)
	}
}

func switchBranch(path string, branch string) ([]byte, error) {
	if !kcgExec.BranchExists(path, branch) {
		return nil, fmt.Errorf("\x1b[31m%s\x1b[0m %s", "invalid branch", branch)
	}

	cmd := exec.Command("git", "switch", branch)
	cmd.Dir = path
	return cmd.CombinedOutput()
}

func convertedBranch(branchArias []string, branch string) string {
	if len(branchArias) == 0 {
		return branch
	}

	rep := regexp.MustCompile(`^[A-Za-z0-9\\\-._]+:[A-Za-z0-9\\\-._]+$`)
	for _, alias := range branchArias {
		if rep.MatchString(alias) {
			if val := strings.Split(alias, ":"); branch == val[0] {
				return val[1]
			}
		}
	}
	return branch
}

func validGroup(groupFlag string, groups []string) bool {
	if groupFlag == "" {
		return true
	}

	for _, group := range groups {
		if groupFlag == group {
			return true
		}
	}

	return false
}

func validFilter(filterFlag string, index string) bool {
	return filterFlag == "" || strings.Contains(index, filterFlag)
}
