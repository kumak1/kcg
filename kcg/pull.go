package kcg

func Pull(config *RepositoryConfig) (string, error) {
	if path, exists := Path(config); exists {
		return kcgGit.Pull(path)
	} else {
		return "", ErrorMessage("invalid path", path)
	}
}
