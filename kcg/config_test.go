package kcg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepositoryConfig_SetAlias(t *testing.T) {
	type args struct {
		aliasList []string
	}
	tests := []struct {
		name  string
		args  args
		wants map[string]string
	}{
		{"valid", args{[]string{"a:b"}}, map[string]string{"a": "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := RepositoryConfig{}
			conf.SetAlias(tt.args.aliasList)
			assert.Equal(t, tt.wants, conf.BranchAlias)
		})
	}
}

func TestRepositoryConfig_AddAlias(t *testing.T) {
	type args struct {
		alias string
	}
	tests := []struct {
		name  string
		args  args
		wants map[string]string
	}{
		{"valid", args{"a:b"}, map[string]string{"a": "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := RepositoryConfig{}
			conf.AddAlias(tt.args.alias)
			assert.Equal(t, tt.wants, conf.BranchAlias)
		})
	}
}
