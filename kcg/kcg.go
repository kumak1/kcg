package kcg

import (
	"fmt"
	"github.com/kumak1/kcg/exec"
	"github.com/kumak1/kcg/exec/ghq"
	"github.com/kumak1/kcg/exec/git"
	"strings"
)

var (
	kcgExec exec.Interface
	kcgGit  git.Interface
	kcgGhq  ghq.Interface
)

func Initialize(config Config) {
	setConfig(config)

	if kcgExec == nil {
		kcgExec = exec.New()
	}
	if kcgGit == nil {
		kcgGit = git.New(kcgExec)
	}
	if kcgGhq == nil {
		kcgGhq = ghq.New(kcgExec)
	}
}

func ValidMessage(colorText string, whiteText string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m %s\n", colorText, whiteText)
}

func ErrorMessage(colorText string, whiteText string) error {
	return fmt.Errorf("\x1b[31m%s\x1b[0m %s\n", colorText, whiteText)
}

func WarningMessage(colorText string, whiteText string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m %s\n", colorText, whiteText)
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
