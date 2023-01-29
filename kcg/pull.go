package kcg

import "github.com/kumak1/kcg/exec/git"

func Pull(config *RepositoryConfig) (string, error) {
	if path, exists := Path(config); exists {
		return git.Pull(path)
	} else {
		return "", ErrorMessage("invalid path", path)
	}
}
