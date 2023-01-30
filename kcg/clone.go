package kcg

import (
	"github.com/kumak1/kcg/exec/ghq"
	"github.com/kumak1/kcg/exec/git"
)

func Clone(config *RepositoryConfig) (string, error) {
	if config.Repo == "" {
		return "", ErrorMessage("error", "repo is empty")
	}

	if path, exists := Path(config); !exists {
		if useGhq {
			return ghq.Get(config.Repo)
		} else {
			return git.Clone(config.Repo, path)
		}
	} else {
		return WarningMessage("exists", path), nil
	}
}
