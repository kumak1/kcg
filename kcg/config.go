package kcg

type Config struct {
	Ghq   bool
	Repos map[string]*RepositoryConfig
}

type RepositoryConfig struct {
	Repo   string
	Path   string
	Groups []string
	Setup  []string
}

func ValidRepo(repoFlag string, index string) bool {
	return repoFlag == "" || repoFlag == index
}

func ValidGroup(groupFlag string, config *RepositoryConfig) bool {
	if groupFlag == "" {
		return true
	}

	for _, group := range config.Groups {
		if groupFlag == group {
			return true
		}
	}

	return false
}
