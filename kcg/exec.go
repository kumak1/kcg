package kcg

func Run(config *RepositoryConfig, command string) (string, error) {
	if path, exists := Path(config); exists {
		return kcgExec.Output(path, "sh", "-c", command)
	} else {
		return "", ErrorMessage("invalid path", path)
	}
}
