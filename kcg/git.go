package kcg

import (
	"fmt"
	kcgExec "github.com/kumak1/kcg/exec"
	kcgGhq "github.com/kumak1/kcg/exec/ghq"
	kcgGit "github.com/kumak1/kcg/exec/git"
	"regexp"
	"strings"
)

var repositoryConfig map[string]*RepositoryConfig

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
	Cleanup(*RepositoryConfig) (string, error)
	Clone(*RepositoryConfig) (string, error)
	CurrentBranch(*RepositoryConfig) string
	List(string, string) map[string]*RepositoryConfig
	Path(*RepositoryConfig) (string, bool)
	Pull(*RepositoryConfig) (string, error)
	Run(*RepositoryConfig, string) (string, error)
	Switch(*RepositoryConfig, string) (string, error)

	// 同じ package 以外から実装を許可しない
	private()
}

type git struct{}

type ghq struct{}

const (
	errorMessageFormat = "  \x1b[31m%s\x1b[0m %s\n"
	warnMessageFormat  = "  \x1b[33m%s\x1b[0m %s\n"
)

func (g git) Cleanup(config *RepositoryConfig) (string, error) {
	if path, exists := g.Path(config); exists {
		return kcgGit.Cleanup(path)
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid path", path)
	}
}

func (g ghq) Cleanup(config *RepositoryConfig) (string, error) {
	if path, exists := g.Path(config); exists {
		return kcgGit.Cleanup(path)
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid path", path)
	}
}

func (g git) Clone(config *RepositoryConfig) (string, error) {
	if config.Repo == "" {
		return "", fmt.Errorf(errorMessageFormat, "error", "repo is empty")
	}

	if path, exists := g.Path(config); exists {
		return fmt.Sprintf(warnMessageFormat, "exists", path), nil
	} else {
		return kcgGit.Clone(config.Repo, path)
	}
}

func (g ghq) Clone(config *RepositoryConfig) (string, error) {
	if config.Repo == "" {
		return "", fmt.Errorf(errorMessageFormat, "error", "repo is empty")
	}

	if path, exists := g.Path(config); exists {
		return fmt.Sprintf(warnMessageFormat, "exists", path), nil
	} else {
		return kcgGhq.Get(config.Repo)
	}
}

func (g git) CurrentBranch(config *RepositoryConfig) string {
	if path, exists := g.Path(config); exists {
		return kcgGit.CurrentBranchName(path)
	} else {
		return ""
	}
}

func (g ghq) CurrentBranch(config *RepositoryConfig) string {
	if path, exists := g.Path(config); exists {
		return kcgGit.CurrentBranchName(path)
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
		path, _ := kcgGhq.Path(config.Repo)
		return path, path != "" && kcgExec.FileExists(path)
	}
}

func (g git) Pull(config *RepositoryConfig) (string, error) {
	if path, exists := g.Path(config); exists {
		return kcgGit.Pull(path)
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid path", path)
	}
}

func (g ghq) Pull(config *RepositoryConfig) (string, error) {
	if path, exists := g.Path(config); exists {
		return kcgGit.Pull(path)
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid path", path)
	}
}

func (g git) Run(config *RepositoryConfig, command string) (string, error) {
	if path, exists := g.Path(config); exists {
		return kcgExec.Output(path, "sh", "-c", command)
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid path", path)
	}
}

func (g ghq) Run(config *RepositoryConfig, command string) (string, error) {
	if path, exists := g.Path(config); exists {
		return kcgExec.Output(path, "sh", "-c", command)
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid path", path)
	}
}

func (g git) Switch(config *RepositoryConfig, branch string) (string, error) {
	if path, exists := g.Path(config); exists {
		return switchBranch(path, convertedBranch(config.Alias, branch))
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid path", path)
	}
}

func (g ghq) Switch(config *RepositoryConfig, branch string) (string, error) {
	if path, exists := g.Path(config); exists {
		return switchBranch(path, convertedBranch(config.Alias, branch))
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid path", path)
	}
}

func switchBranch(path string, branch string) (string, error) {
	if kcgGit.BranchExists(path, branch) {
		return kcgGit.Switch(path, branch)
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid branch", branch)
	}
}

func (g git) private() {}
func (g ghq) private() {}

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
