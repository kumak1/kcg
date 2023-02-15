package pkg

func Path(config *RepositoryConfig) (string, bool) {
	if useGhq {
		if config.Path != "" {
			return config.Path, kcgExec.FileExists(config.Path)
		} else {
			if config.Repo == "" {
				return "", false
			}
			path, _ := kcgGhq.Path(config.Repo)
			return path, path != "" && kcgExec.FileExists(path)
		}
	} else {
		return config.Path, config.Path != "" && kcgExec.FileExists(config.Path)
	}
}
