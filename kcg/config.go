package kcg

type Config struct {
	Ghq   bool
	Repos map[string]*RepositoryConfig
}

type RepositoryConfig struct {
	Repo   string
	Path   string
	Alias  []string
	Group  []string
	Setup  []string
	Update []string
}

var useGhq bool
var repositoryConfig map[string]*RepositoryConfig

func SetConfig(config Config) {
	useGhq = config.Ghq
	repositoryConfig = config.Repos
}
