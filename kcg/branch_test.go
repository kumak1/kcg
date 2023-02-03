package kcg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrentBranch(t *testing.T) {
	f := func() *RepositoryConfig {
		var a = &RepositoryConfig{}
		a.Path = "path"
		a.Repo = "repo"
		return a
	}
	type args struct {
		config      *RepositoryConfig
		mockUseGhq  bool
		mockExists  bool
		mockPath    string
		mockPathErr bool
		mockBranch  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"valid", args{f(), true, true, "", true, "branch"}, "branch"},
		{"valid", args{f(), true, false, "", true, "branch"}, ""},
	}
	for _, tt := range tests {
		useGhq = tt.args.mockUseGhq

		testExecObj := new(MockedExecInterface)
		testExecObj.On("FileExists").Return(tt.args.mockExists)
		kcgExec = testExecObj

		testGitObj := new(MockedGitInterface)
		testGitObj.On("CurrentBranchName").Return(tt.args.mockBranch)
		kcgGit = testGitObj

		testGhqObj := new(MockedGhqInterface)
		testGhqObj.On("Path").Return(tt.args.mockPath, tt.args.mockPathErr)
		kcgGhq = testGhqObj

		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CurrentBranch(tt.args.config), "CurrentBranch(%v)", tt.args.config)
		})
	}
}
