package kcg

import "github.com/kumak1/kcg/exec/git"

func CurrentBranch(config *RepositoryConfig) string {
	if path, exists := Path(config); exists {
		return git.CurrentBranchName(path)
	} else {
		return ""
	}
}
