package pkg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	f := func(repo string, path string) *RepositoryConfig {
		var a = &RepositoryConfig{}
		a.Path = path
		a.Repo = repo
		return a
	}
	type args struct {
		config        *RepositoryConfig
		command       string
		mockUseGhq    bool
		mockExists    bool
		mockPath      string
		mockPathErr   bool
		mockOutput    string
		mockOutputErr bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{"valid", args{f("repo", "path"), "command", true, true, "path", true, "cleanup", true}, "cleanup", assert.NoError},
		{"invalid", args{f("", ""), "command", true, true, "", true, "", true}, "", assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useGhq = tt.args.mockUseGhq

			testExecObj := new(MockedExecInterface)
			testExecObj.On("FileExists").Return(tt.args.mockExists)
			testExecObj.On("Output").Return(tt.args.mockOutput, tt.args.mockOutputErr)
			kcgExec = testExecObj

			testGitObj := new(MockedGitInterface)
			kcgGit = testGitObj

			testGhqObj := new(MockedGhqInterface)
			testGhqObj.On("Path").Return(tt.args.mockPath, tt.args.mockPathErr)
			kcgGhq = testGhqObj

			got, err := Run(tt.args.config, tt.args.command)
			if !tt.wantErr(t, err, fmt.Sprintf("Run(%v, %v)", tt.args.config, tt.args.command)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Run(%v, %v)", tt.args.config, tt.args.command)
		})
	}
}
