package kcg

import "github.com/kumak1/kcg/exec"

func Run(config *RepositoryConfig, command string) (string, error) {
	if path, exists := Path(config); exists {
		return exec.New().Output(path, "sh", "-c", command)
	} else {
		return "", ErrorMessage("invalid path", path)
	}
}
