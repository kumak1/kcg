package kcg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClone(t *testing.T) {
	f := func(repo string, path string) *RepositoryConfig {
		var a = &RepositoryConfig{}
		a.Path = path
		a.Repo = repo
		return a
	}
	type args struct {
		config       *RepositoryConfig
		mockUseGhq   bool
		mockExists   bool
		mockPath     string
		mockPathErr  bool
		mockClone    string
		mockCloneErr bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{"empty repo", args{f("", ""), true, true, "", true, "", true}, "", assert.Error},
		{"invalid path", args{f("repo", "path"), false, true, "path", true, "", true}, "\x1b[33mexists\x1b[0m path\n", assert.NoError},
		{"ghq get", args{f("repo", ""), true, true, "", true, "a", true}, "a", assert.NoError},
		{"git clone", args{f("repo", ""), false, true, "", true, "a", true}, "a", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useGhq = tt.args.mockUseGhq

			testExecObj := new(MockedExecInterface)
			testExecObj.On("FileExists").Return(tt.args.mockExists)
			kcgExec = testExecObj

			testGitObj := new(MockedGitInterface)
			testGitObj.On("Clone").Return(tt.args.mockClone, tt.args.mockCloneErr)
			kcgGit = testGitObj

			testGhqObj := new(MockedGhqInterface)
			testGhqObj.On("Get").Return(tt.args.mockClone, tt.args.mockCloneErr)
			testGhqObj.On("Path").Return(tt.args.mockPath, tt.args.mockPathErr)
			kcgGhq = testGhqObj

			got, err := Clone(tt.args.config)
			if !tt.wantErr(t, err, fmt.Sprintf("Clone(%v)", tt.args.config)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Clone(%v)", tt.args.config)
		})
	}
}
