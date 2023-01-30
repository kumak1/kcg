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

var (
	useGhq           bool
	repositoryConfig map[string]*RepositoryConfig
)

func setConfig(config Config) {
	useGhq = config.Ghq
	repositoryConfig = config.Repos
}
