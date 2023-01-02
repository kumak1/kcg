package kcg

type Config struct {
	Repos map[string]*RepositoryConfig
}

type RepositoryConfig struct {
	Repo          string
	Path          string
	SetupCommands []string
}
