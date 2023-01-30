package kcg

func Clone(config *RepositoryConfig) (string, error) {
	if config.Repo == "" {
		return "", ErrorMessage("error", "repo is empty")
	}

	if path, exists := Path(config); !exists {
		if useGhq {
			return kcgGhq.Get(config.Repo)
		} else {
			return kcgGit.Clone(config.Repo, path)
		}
	} else {
		return WarningMessage("exists", path), nil
	}
}
