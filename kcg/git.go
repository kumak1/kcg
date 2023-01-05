package kcg

import (
	"fmt"
	kcgExec "github.com/kumak1/kcg/kcg/exec"
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
	Cleanup(*RepositoryConfig) error
	Clone(*RepositoryConfig) error
	Path(*RepositoryConfig) string
	Pull(*RepositoryConfig) error
	Setup(*RepositoryConfig) error
	Switch(*RepositoryConfig, string) error
}

type git struct{}

type ghq struct{}

func (g git) Cleanup(config *RepositoryConfig) error {
	return cleanup(config.Path)
}

func (g ghq) Cleanup(config *RepositoryConfig) error {
	if path, err := kcgExec.GhqPath(config.Repo); err != nil {
		return err
	} else {
		return cleanup(path)
	}
}

func cleanup(path string) error {
	fmt.Println(path)
	cmd := exec.Command("sh", "-c", "git branch --merged|egrep -v '\\*|develop|main|master'|xargs git branch -d")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (g git) Clone(config *RepositoryConfig) error {
	if kcgExec.DirExists(config.Path) {
		fmt.Println("exists: " + config.Path)
		return nil
	}

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

func (g git) Path(config *RepositoryConfig) string {
	return config.Path
}

func (g ghq) Path(config *RepositoryConfig) string {
	path, _ := kcgExec.GhqPath(config.Repo)
	return path
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
