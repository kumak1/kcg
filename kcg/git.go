package kcg

import (
	"fmt"
	kcgExec "github.com/kumak1/kcg/exec"
	kcgGhq "github.com/kumak1/kcg/exec/ghq"
	kcgGit "github.com/kumak1/kcg/exec/git"
	"regexp"
	"strings"
)

var useGhq bool
var repositoryConfig map[string]*RepositoryConfig

func SetConfig(config Config) {
	useGhq = config.Ghq
	repositoryConfig = config.Repos
}

func Cleanup(config *RepositoryConfig) (string, error) {
	if path, exists := Path(config); exists {
		return kcgGit.Cleanup(path)
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid path", path)
	}
}

func Clone(config *RepositoryConfig) (string, error) {
	if config.Repo == "" {
		return "", fmt.Errorf(errorMessageFormat, "error", "repo is empty")
	}

	if path, exists := Path(config); exists {
		return fmt.Sprintf(warnMessageFormat, "exists", path), nil
	} else {
		return kcgGit.Clone(config.Repo, path)
	}
}

func CurrentBranch(config *RepositoryConfig) string {
	if path, exists := Path(config); exists {
		return kcgGit.CurrentBranchName(path)
	} else {
		return ""
	}
}

func List(group string, filter string) map[string]*RepositoryConfig {
	repositoryConfigs := map[string]*RepositoryConfig{}
	for index, repo := range repositoryConfig {
		if validGroup(group, repo.Group) && validFilter(filter, index) {
			repositoryConfigs[index] = repo
		}
	}
	return repositoryConfigs
}

func Pull(config *RepositoryConfig) (string, error) {
	if path, exists := Path(config); exists {
		return kcgGit.Pull(path)
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid path", path)
	}
}

func Path(config *RepositoryConfig) (string, bool) {
	if useGhq {
		if config.Path != "" {
			return config.Path, kcgExec.FileExists(config.Path)
		} else {
			if config.Repo == "" {
				return "", false
			}
			path, _ := kcgGhq.Path(config.Repo)
			return path, path != "" && kcgExec.FileExists(path)
		}
	} else {
		return config.Path, config.Path != "" && kcgExec.FileExists(config.Path)
	}
}

func Run(config *RepositoryConfig, command string) (string, error) {
	if path, exists := Path(config); exists {
		return kcgExec.Output(path, "sh", "-c", command)
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid path", path)
	}
}

func Switch(config *RepositoryConfig, branch string) (string, error) {
	if path, exists := Path(config); exists {
		convertedBranch := convertedBranch(config.Alias, branch)
		if kcgGit.BranchExists(path, convertedBranch) {
			return kcgGit.Switch(path, convertedBranch)
		} else {
			return "", fmt.Errorf(errorMessageFormat, "invalid branch", convertedBranch)
		}
	} else {
		return "", fmt.Errorf(errorMessageFormat, "invalid path", path)
	}
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
