package exec

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{"command.go"}, true},
		{"invalid", args{"invalid"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, New().FileExists(tt.args.path), "fileExists(%v)", tt.args.path)
		})
	}
}

func TestDirExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{"."}, true},
		{"invalid", args{"invalid"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, New().DirExists(tt.args.path), "dirExists(%v)", tt.args.path)
		})
	}
}

func TestOutput(t *testing.T) {
	type args struct {
		path string
		name string
		arg  []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{"valid", args{"", "echo", []string{"a"}}, "a", assert.NoError},
		{"invalid", args{"", "invalid", []string{""}}, "", assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New().Output(tt.args.path, tt.args.name, tt.args.arg...)
			if !tt.wantErr(t, err, fmt.Sprintf("output(%v, %v, %v)", tt.args.path, tt.args.name, tt.args.arg)) {
				return
			}
			assert.Equalf(t, tt.want, got, "output(%v, %v, %v)", tt.args.path, tt.args.name, tt.args.arg)
		})
	}
}

func TestNotError(t *testing.T) {
	type args struct {
		path string
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{"", "ls"}, true},
		{"invalid", args{"", "invalid"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, New().NotError(tt.args.path, tt.args.name), "notError(%v, %v, %v)", tt.args.path, tt.args.name)
		})
	}
}
