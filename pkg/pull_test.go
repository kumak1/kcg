package pkg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPull(t *testing.T) {
	f := func() *RepositoryConfig {
		var a = &RepositoryConfig{}
		a.Path = "path"
		a.Repo = "repo"
		return a
	}
	type args struct {
		config      *RepositoryConfig
		mockPull    string
		mockPullErr bool
		mockPath    string
		mockPathErr bool
		mockExists  bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{"valid", args{f(), "valid", true, "valid", true, true}, "valid", assert.NoError},
		{"invalid", args{f(), "invalid", true, "invalid", true, false}, "", assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useGhq = false

			testExecObj := new(MockedExecInterface)
			testExecObj.On("FileExists").Return(tt.args.mockExists)
			kcgExec = testExecObj

			testGitObj := new(MockedGitInterface)
			testGitObj.On("Pull").Return(tt.args.mockPath, tt.args.mockPullErr)
			kcgGit = testGitObj

			testGhqObj := new(MockedGhqInterface)
			testGhqObj.On("Path").Return(tt.args.mockPath, tt.args.mockPathErr)
			kcgGhq = testGhqObj

			got, err := Pull(tt.args.config)
			if !tt.wantErr(t, err, fmt.Sprintf("Pull(%v)", tt.args.config)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Pull(%v)", tt.args.config)
		})
	}
}
