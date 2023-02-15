package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
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

func TestListParallelFor(t *testing.T) {
	type args struct {
		fn     func(key string, repoConf *RepositoryConfig)
		group  string
		filter string
		config map[string]*RepositoryConfig
	}
	aConfig := map[string]*RepositoryConfig{"a": {Group: []string{"group"}}}
	tests := []struct {
		name string
		args args
	}{
		{"valid", args{
			func(key string, repoConf *RepositoryConfig) {

			}, "", "", aConfig}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListParallelFor(tt.args.fn, tt.args.group, tt.args.filter)
		})
	}
}
