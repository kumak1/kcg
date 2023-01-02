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
	fmt.Println(config.Path)
	cmd := exec.Command("git", "pull")
	cmd.Dir = config.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (g ghq) Pull(config *RepositoryConfig) error {
	fmt.Println(config.Path)
	cmd := exec.Command("git", "pull")
	cmd.Dir = config.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func ghqPath() ([]byte, error) {
	cmd := exec.Command("ghq", "list", "-p", "-e", "kumaoche")
	return cmd.Output()
}
