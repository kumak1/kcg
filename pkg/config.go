package pkg

import (
	"regexp"
	"strings"
)

type Config struct {
	Ghq   bool
	Repos map[string]*RepositoryConfig
}

type RepositoryConfig struct {
	Repo        string
	Path        string
	BranchAlias map[string]string
	Group       []string
	Exec        map[string][]string
}

func (conf *RepositoryConfig) AddAlias(alias string) {
	if conf.BranchAlias == nil {
		conf.BranchAlias = map[string]string{}
	}

	rep := regexp.MustCompile(`^[A-Za-z0-9\\\-._]+:[A-Za-z0-9\\\-._]+$`)
	if rep.MatchString(alias) {
		val := strings.Split(alias, ":")
		conf.BranchAlias[val[0]] = val[1]
	}
}

func (conf *RepositoryConfig) SetAlias(aliasList []string) {
	if conf.BranchAlias == nil {
		conf.BranchAlias = map[string]string{}
	}

	rep := regexp.MustCompile(`^[A-Za-z0-9\\\-._]+:[A-Za-z0-9\\\-._]+$`)
	for _, alias := range aliasList {
		if rep.MatchString(alias) {
			val := strings.Split(alias, ":")
			conf.BranchAlias[val[0]] = val[1]
		}
	}
}

var (
	useGhq           bool
	repositoryConfig map[string]*RepositoryConfig
)

func setConfig(config Config) {
	useGhq = config.Ghq
	repositoryConfig = config.Repos
}
