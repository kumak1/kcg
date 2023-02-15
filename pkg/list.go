package pkg

import "sync"

func List(group string, filter string) map[string]*RepositoryConfig {
	configs := map[string]*RepositoryConfig{}
	for index, repo := range repositoryConfig {
		if validGroup(group, repo.Group) && validFilter(filter, index) {
			configs[index] = repo
		}
	}
	return configs
}

func ListParallelFor(fn func(key string, repoConf *RepositoryConfig), group string, filter string) {
	var wg sync.WaitGroup
	for key, conf := range List(group, filter) {
		wg.Add(1)
		key := key
		conf := conf
		go func() {
			fn(key, conf)
			wg.Done()
		}()
	}
	wg.Wait()
}
