package kcg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type (
	MockedExecInterface struct {
		mock.Mock
	}
	MockedGitInterface struct {
		mock.Mock
	}
	MockedGhqInterface struct {
		mock.Mock
	}
)

func (m *MockedExecInterface) FileExists(s string) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockedExecInterface) DirExists(s string) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockedExecInterface) Output(s string, s2 string, s3 ...string) (string, error) {
	args := m.Called()
	if args.Bool(1) {
		return args.String(0), nil
	} else {
		return args.String(0), fmt.Errorf("err")
	}
}

func (m *MockedExecInterface) NotError(s string, s2 string, s3 ...string) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockedGitInterface) BranchExists(path string, branch string) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockedGitInterface) CurrentBranchName(path string) string {
	args := m.Called()
	return args.String(0)
}
func (m *MockedGitInterface) Switch(path string, branch string) (string, error) {
	args := m.Called()
	if args.Bool(1) {
		return args.String(0), nil
	} else {
		return args.String(0), fmt.Errorf("err")
	}
}
func (m *MockedGitInterface) Pull(path string) (string, error) {
	args := m.Called()
	if args.Bool(1) {
		return args.String(0), nil
	} else {
		return args.String(0), fmt.Errorf("err")
	}
}
func (m *MockedGitInterface) Clone(repo string, path string) (string, error) {
	args := m.Called()
	if args.Bool(1) {
		return args.String(0), nil
	} else {
		return args.String(0), fmt.Errorf("err")
	}
}
func (m *MockedGitInterface) Cleanup(path string) (string, error) {
	args := m.Called()
	if args.Bool(1) {
		return args.String(0), nil
	} else {
		return args.String(0), fmt.Errorf("err")
	}
}
func (m *MockedGitInterface) OriginUrl(path string) (string, error) {
	args := m.Called()
	if args.Bool(1) {
		return args.String(0), nil
	} else {
		return args.String(0), fmt.Errorf("err")
	}
}

func (m *MockedGhqInterface) Valid() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockedGhqInterface) Get(s string) (string, error) {
	args := m.Called()
	if args.Bool(1) {
		return args.String(0), nil
	} else {
		return args.String(0), fmt.Errorf("err")
	}
}

func (m *MockedGhqInterface) Path(s string) (string, error) {
	args := m.Called()
	if args.Bool(1) {
		return args.String(0), nil
	} else {
		return args.String(0), fmt.Errorf("err")
	}
}

func (m *MockedGhqInterface) List() []string {
	args := m.Called()
	return []string{args.String(0)}
}

func TestValidMessage(t *testing.T) {
	type args struct {
		colorText string
		whiteText string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"get message", args{"check", "text"}, "\x1b[32mcheck\x1b[0m text\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ValidMessage(tt.args.colorText, tt.args.whiteText), "ValidMessage(%v, %v)", tt.args.colorText, tt.args.whiteText)
		})
	}
}

func TestErrorMessage(t *testing.T) {
	type args struct {
		colorText string
		whiteText string
	}

	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{"get message", args{"check", "text"}, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, ErrorMessage(tt.args.colorText, tt.args.whiteText), fmt.Sprintf("ErrorMessage(%v, %v)", tt.args.colorText, tt.args.whiteText))
		})
	}
}

func TestWarningMessage(t *testing.T) {
	type args struct {
		colorText string
		whiteText string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"get message", args{"check", "text"}, "\x1b[33mcheck\x1b[0m text\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, WarningMessage(tt.args.colorText, tt.args.whiteText), "WarningMessage(%v, %v)", tt.args.colorText, tt.args.whiteText)
		})
	}
}

func Test_validGroup(t *testing.T) {
	type args struct {
		groupFlag string
		groups    []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid 1", args{"", []string{"a", "b"}}, true},
		{"valid 2", args{"a", []string{"a", "b"}}, true},
		{"valid 3", args{"b", []string{"a", "b"}}, true},
		{"invalid", args{"c", []string{"a", "b"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, validGroup(tt.args.groupFlag, tt.args.groups), "validGroup(%v, %v)", tt.args.groupFlag, tt.args.groups)
		})
	}
}

func Test_validFilter(t *testing.T) {
	type args struct {
		filterFlag string
		index      string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"no filter", args{"", ""}, true},
		{"valid filter", args{"a", "valid"}, true},
		{"invalid filter", args{"b", "valid"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, validFilter(tt.args.filterFlag, tt.args.index), "validFilter(%v, %v)", tt.args.filterFlag, tt.args.index)
		})
	}
}
