package kcg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanup(t *testing.T) {
	f := func(repo string, path string) *RepositoryConfig {
		var a = &RepositoryConfig{}
		a.Path = path
		a.Repo = repo
		return a
	}
	type args struct {
		config         *RepositoryConfig
		mockUseGhq     bool
		mockExists     bool
		mockPath       string
		mockPathErr    bool
		mockCleanup    string
		mockCleanupErr bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{"valid", args{f("repo", "path"), true, true, "path", true, "cleanup", true}, "cleanup", assert.NoError},
		{"invalid", args{f("", ""), true, true, "", true, "", true}, "", assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useGhq = tt.args.mockUseGhq

			testExecObj := new(MockedExecInterface)
			testExecObj.On("FileExists").Return(tt.args.mockExists)
			kcgExec = testExecObj

			testGitObj := new(MockedGitInterface)
			testGitObj.On("Cleanup").Return(tt.args.mockCleanup, tt.args.mockCleanupErr)
			kcgGit = testGitObj

			testGhqObj := new(MockedGhqInterface)
			testGhqObj.On("Path").Return(tt.args.mockPath, tt.args.mockPathErr)
			kcgGhq = testGhqObj

			got, err := Cleanup(tt.args.config)
			if !tt.wantErr(t, err, fmt.Sprintf("Cleanup(%v)", tt.args.config)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Cleanup(%v)", tt.args.config)
		})
	}
}
