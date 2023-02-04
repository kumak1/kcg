package kcg

func List(group string, filter string) map[string]*RepositoryConfig {
	configs := map[string]*RepositoryConfig{}
	for index, repo := range repositoryConfig {
		if validGroup(group, repo.Group) && validFilter(filter, index) {
			configs[index] = repo
		}
	}
	return configs
}
