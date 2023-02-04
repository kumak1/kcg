package kcg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwitch(t *testing.T) {
	f := func(repo string, path string) *RepositoryConfig {
		var a = &RepositoryConfig{}
		a.Path = path
		a.Repo = repo
		a.Alias = []string{"a:b"}
		return a
	}
	type args struct {
		config           *RepositoryConfig
		branch           string
		mockUseGhq       bool
		mockExists       bool
		mockBranchExists bool
		mockPath         string
		mockPathErr      bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{"invalid path", args{f("", ""), "branch", false, false, false, "", true}, "", assert.Error},
		{"invalid branch", args{f("repo", "path"), "branch", false, true, false, "", true}, "", assert.Error},
		{"valid", args{f("repo", "path"), "branch", false, true, true, "", true}, "valid", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useGhq = tt.args.mockUseGhq

			testExecObj := new(MockedExecInterface)
			testExecObj.On("FileExists").Return(tt.args.mockExists)
			kcgExec = testExecObj

			testGitObj := new(MockedGitInterface)
			testGitObj.On("BranchExists").Return(tt.args.mockBranchExists)
			testGitObj.On("Switch").Return("valid", true)
			kcgGit = testGitObj

			testGhqObj := new(MockedGhqInterface)
			testGhqObj.On("Path").Return(tt.args.mockPath, tt.args.mockPathErr)
			kcgGhq = testGhqObj

			got, err := Switch(tt.args.config, tt.args.branch)
			if !tt.wantErr(t, err, fmt.Sprintf("Switch(%v, %v)", tt.args.config, tt.args.branch)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Switch(%v, %v)", tt.args.config, tt.args.branch)
		})
	}
}

func Test_convertedBranch(t *testing.T) {
	type args struct {
		branchArias []string
		branch      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"arias empty", args{[]string{}, "branch"}, "branch"},
		{"arias not hit", args{[]string{"a:b"}, "branch"}, "branch"},
		{"arias hit", args{[]string{"a:b"}, "a"}, "b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, convertedBranch(tt.args.branchArias, tt.args.branch), "convertedBranch(%v, %v)", tt.args.branchArias, tt.args.branch)
		})
	}
}
