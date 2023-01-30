package kcg

import (
	"regexp"
	"strings"
)

func Switch(config *RepositoryConfig, branch string) (string, error) {
	if path, exists := Path(config); exists {
		convertedBranch := convertedBranch(config.Alias, branch)
		if kcgGit.BranchExists(path, convertedBranch) {
			return kcgGit.Switch(path, convertedBranch)
		} else {
			return "", ErrorMessage("invalid branch", convertedBranch)
		}
	} else {
		return "", ErrorMessage("invalid path", path)
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
