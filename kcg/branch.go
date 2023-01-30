package kcg

func CurrentBranch(config *RepositoryConfig) string {
	if path, exists := Path(config); exists {
		return kcgGit.CurrentBranchName(path)
	} else {
		return ""
	}
}
