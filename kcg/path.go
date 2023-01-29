package kcg

import (
	"github.com/kumak1/kcg/exec"
	"github.com/kumak1/kcg/exec/ghq"
)

func Path(config *RepositoryConfig) (string, bool) {
	if useGhq {
		if config.Path != "" {
			return config.Path, exec.FileExists(config.Path)
		} else {
			if config.Repo == "" {
				return "", false
			}
			path, _ := ghq.Path(config.Repo)
			return path, path != "" && exec.FileExists(path)
		}
	} else {
		return config.Path, config.Path != "" && exec.FileExists(config.Path)
	}
}
