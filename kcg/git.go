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
	return config.Path, kcgExec.FileExists(config.Path)
}

func (g ghq) Path(config *RepositoryConfig) (string, bool) {
	if config.Path != "" {
		return config.Path, kcgExec.FileExists(config.Path)
	} else {
		path, _ := kcgExec.GhqPath(config.Repo)
		return path, kcgExec.FileExists(path)
	}
}

func (g git) Pull(config *RepositoryConfig) error {
	return pull(config.Path)
}

func (g ghq) Pull(config *RepositoryConfig) error {
	if path, err := kcgExec.GhqPath(config.Repo); err != nil {
		return err
	} else {
		return pull(path)
	}
}

func pull(path string) error {
	fmt.Println(path)
	cmd := exec.Command("git", "pull")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (g git) Switch(config *RepositoryConfig, branch string) error {
	return switchBranch(config.Path, branch)
}

func (g ghq) Switch(config *RepositoryConfig, branch string) error {
	if path, err := kcgExec.GhqPath(config.Repo); err != nil {
		return err
	} else {
		return switchBranch(path, branch)
	}
}

func switchBranch(path string, branch string) error {
	fmt.Println(path)
	if !kcgExec.BranchExists(path, branch) {
		fmt.Println("'" + branch + "' branch is not exists.")
		return nil
	}

	cmd := exec.Command("git", "switch", branch)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (g git) Setup(config *RepositoryConfig) error {
	return setup(config.Path, config.Setup)
}

func (g ghq) Setup(config *RepositoryConfig) error {
	if path, err := kcgExec.GhqPath(config.Repo); err != nil {
		return err
	} else {
		return setup(path, config.Setup)
	}
}

func setup(path string, commands []string) error {
	fmt.Println("\n" + path)

	if len(commands) == 0 {
		fmt.Println("no settings.")
	}

	for _, setupCommand := range commands {
		fmt.Println("run: " + setupCommand)
		cmd := exec.Command("sh", "-c", setupCommand)
		cmd.Dir = path
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
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
