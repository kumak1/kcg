package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPath(t *testing.T) {
	f := func(path string, repo string) *RepositoryConfig {
		var a = &RepositoryConfig{}
		a.Path = path
		a.Repo = repo
		return a
	}
	type args struct {
		config      *RepositoryConfig
		mockUseGhq  bool
		mockExists  bool
		mockPath    string
		mockPathErr bool
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{"ghq config path", args{f("valid", ""), true, true, "", true}, "valid", true},
		{"ghq empty repo", args{f("", ""), true, true, "empty", true}, "", false},
		{"ghq fill repo", args{f("", "repo"), true, true, "valid", true}, "valid", true},
		{"git config path", args{f("path", ""), false, true, "", true}, "path", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useGhq = tt.args.mockUseGhq

			testExecObj := new(MockedExecInterface)
			testExecObj.On("FileExists").Return(tt.args.mockExists)
			kcgExec = testExecObj

			testGitObj := new(MockedGitInterface)
			kcgGit = testGitObj

			testGhqObj := new(MockedGhqInterface)
			testGhqObj.On("Path").Return(tt.args.mockPath, tt.args.mockPathErr)
			kcgGhq = testGhqObj

			got, got1 := Path(tt.args.config)
			assert.Equalf(t, tt.want, got, "Path(%v)", tt.args.config)
			assert.Equalf(t, tt.want1, got1, "Path(%v)", tt.args.config)
		})
	}
}
