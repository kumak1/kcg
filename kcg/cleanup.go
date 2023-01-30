package kcg

import (
	"github.com/kumak1/kcg/exec/git"
)

func Cleanup(config *RepositoryConfig) (string, error) {
	if path, exists := Path(config); exists {
		return git.Cleanup(path)
	} else {
		return "", ErrorMessage("invalid path", path)
	}
}
