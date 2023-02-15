package pkg

func Cleanup(config *RepositoryConfig) (string, error) {
	if path, exists := Path(config); exists {
		return kcgGit.Cleanup(path)
	} else {
		return "", ErrorMessage("invalid path", path)
	}
}
