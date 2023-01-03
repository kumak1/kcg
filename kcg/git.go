package kcg

import (
	"fmt"
	"os"
	"os/exec"
)

func GitCommand(config Config) GitOperateInterface {
	var gitOperator GitOperateInterface

	if config.Ghq {
		gitOperator = ghq{}
	} else {
		gitOperator = git{}
	}

	return gitOperator
}

type GitOperateInterface interface {
	Clone(*RepositoryConfig) error
	Pull(*RepositoryConfig) error
	Switch(*RepositoryConfig, string) error
}

type git struct{}

type ghq struct{}

func (g git) Clone(config *RepositoryConfig) error {
	cmd := exec.Command("git", "clone", config.Repo, config.Path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (g ghq) Clone(config *RepositoryConfig) error {
	cmd := exec.Command("ghq", "get", config.Repo)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (g git) Pull(config *RepositoryConfig) error {
	return pull(config.Path)
}

func (g ghq) Pull(config *RepositoryConfig) error {
	return pull(config.Path)
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
	return switchBranch(config.Path, branch)
}

func switchBranch(path string, branch string) error {
	fmt.Println("\n" + path)
	cmd := exec.Command("git", "switch", branch)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
