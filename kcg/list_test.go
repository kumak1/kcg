package kcg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	//f := func(repo string, path string) *RepositoryConfig {
	//	var a = &RepositoryConfig{}
	//	a.Path = path
	//	a.Repo = repo
	//	a.Alias = []string{"a:b"}
	//	return a
	//}
	emptyConfig := map[string]*RepositoryConfig{}
	aConfig := map[string]*RepositoryConfig{"a": {Group: []string{"group"}}}
	type args struct {
		group  string
		filter string
		config map[string]*RepositoryConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]*RepositoryConfig
	}{
		{"emptyConfig", args{"", "", emptyConfig}, emptyConfig},
		{"valid", args{"", "a", aConfig}, aConfig},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repositoryConfig = tt.args.config
			assert.Equalf(t, tt.want, List(tt.args.group, tt.args.filter), "List(%v, %v)", tt.args.group, tt.args.filter)
		})
	}
}
