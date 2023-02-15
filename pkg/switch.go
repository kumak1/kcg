package pkg

func Switch(config *RepositoryConfig, branch string) (string, error) {
	if path, exists := Path(config); exists {
		convertedBranch := convertedBranch(config.BranchAlias, branch)
		if kcgGit.BranchExists(path, convertedBranch) {
			return kcgGit.Switch(path, convertedBranch)
		} else {
			return "", ErrorMessage("invalid branch", convertedBranch)
		}
	} else {
		return "", ErrorMessage("invalid path", path)
	}
}

func convertedBranch(branchArias map[string]string, branch string) string {
	if len(branchArias) == 0 {
		return branch
	}

	for key, alias := range branchArias {
		if branch == key {
			return alias
		}
	}
	return branch
}
