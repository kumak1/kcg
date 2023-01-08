package kcg

import (
	"fmt"
	kcgExec "github.com/kumak1/kcg/kcg/exec"
	"os"
	"os/exec"
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
	Cleanup(*RepositoryConfig) error
	Clone(*RepositoryConfig) error
	List(string, string) map[string]*RepositoryConfig
	Path(*RepositoryConfig) (string, bool)
	Pull(*RepositoryConfig) error
	Setup(*RepositoryConfig) error
	Switch(*RepositoryConfig, string) error
}

type git struct{}

type ghq struct{}

func (g git) Cleanup(config *RepositoryConfig) error {
	if path, exists := g.Path(config); exists {
		return cleanup(path)
	} else {
		return fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "not exists", path)
	}
}

func (g ghq) Cleanup(config *RepositoryConfig) error {
	if path, exists := g.Path(config); exists {
		return cleanup(path)
	} else {
		return fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "not exists", path)
	}
}

func cleanup(path string) error {
	cmd := exec.Command("sh", "-c", "git branch --merged|egrep -v '\\*|develop|main|master'|xargs git branch -d")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (g git) Clone(config *RepositoryConfig) error {
	if config.Repo == "" {
		return fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "error", "repo is empty")
	}

	if path, exists := g.Path(config); exists {
		return fmt.Errorf("    \x1b[33m%s\x1b[0m %s", "exists", path)
	} else {
		cmd := exec.Command("git", "clone", config.Repo, path)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
}

func (g ghq) Clone(config *RepositoryConfig) error {
	if config.Repo == "" {
		return fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "error", "repo is empty")
	}

	cmd := exec.Command("ghq", "get", config.Repo)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
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
		if validGroup(group, repo.Groups) && validFilter(filter, index) {
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
		path, _ := kcgExec.GhqPath(config.Repo)
		return path, path != "" && kcgExec.FileExists(path)
	}
}

func (g git) Pull(config *RepositoryConfig) error {
	if path, exists := g.Path(config); exists {
		return pull(path)
	} else {
		return fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "not exists", path)
	}
}

func (g ghq) Pull(config *RepositoryConfig) error {
	if path, exists := g.Path(config); exists {
		return pull(path)
	} else {
		return fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "not exists", path)
	}
}

func pull(path string) error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (g git) Switch(config *RepositoryConfig, branch string) error {
	if path, exists := g.Path(config); exists {
		return switchBranch(path, branch)
	} else {
		return fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "not exists", path)
	}
}

func (g ghq) Switch(config *RepositoryConfig, branch string) error {
	if path, exists := g.Path(config); exists {
		return switchBranch(path, branch)
	} else {
		return fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "not exists", path)
	}
}

func switchBranch(path string, branch string) error {
	if !kcgExec.BranchExists(path, branch) {
		return fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "not exists", branch)
	}

	cmd := exec.Command("git", "switch", branch)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (g git) Setup(config *RepositoryConfig) error {
	if path, exists := g.Path(config); exists {
		return setup(path, config.Setup)
	} else {
		return fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "not exists", path)
	}
}

func (g ghq) Setup(config *RepositoryConfig) error {
	if path, exists := g.Path(config); exists {
		return setup(path, config.Setup)
	} else {
		return fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "not exists", path)
	}
}

func setup(path string, commands []string) error {
	if len(commands) == 0 {
		return fmt.Errorf("    \x1b[33m%s\x1b[0m", "not exists")
	}

	for _, setupCommand := range commands {
		fmt.Printf("    \x1b[32m%s\x1b[0m %s\n", "setup", setupCommand)
		cmd := exec.Command("sh", "-c", setupCommand)
		cmd.Dir = path
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
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
