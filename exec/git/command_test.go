package git

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type (
	MockedInterface struct {
		mock.Mock
	}
	outputArgs struct {
		text string
		err  bool
	}
)

func (m *MockedInterface) FileExists(s string) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockedInterface) DirExists(s string) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockedInterface) Output(s string, s2 string, s3 ...string) (string, error) {
	args := m.Called()
	if args.Bool(1) {
		return args.String(0), nil
	} else {
		return args.String(0), fmt.Errorf("err")
	}
}

func (m *MockedInterface) NotError(s string, s2 string, s3 ...string) bool {
	args := m.Called()
	return args.Bool(0)
}

func Test_defaultExec_BranchExists(t *testing.T) {
	type args struct {
		path   string
		branch string
	}
	tests := []struct {
		name string
		args args
		mock bool
		want bool
	}{
		{"valid", args{"path", "branch"}, true, true},
		{"invalid", args{"path", "branch"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testObj := new(MockedInterface)
			testObj.On("NotError").Return(tt.mock)
			d := New(testObj)
			assert.Equalf(t, tt.want, d.BranchExists(tt.args.path, tt.args.branch), "BranchExists(%v, %v)", tt.args.path, tt.args.branch)

			testObj.AssertExpectations(t)
		})
	}
}

func Test_defaultExec_CurrentBranchName(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name     string
		args     args
		mockArgs outputArgs
		want     string
	}{
		{"valid", args{"path"}, outputArgs{"main", true}, "main"},
		{"invalid", args{"path"}, outputArgs{"main", false}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testObj := new(MockedInterface)
			testObj.On("Output").Return(tt.mockArgs.text, tt.mockArgs.err)
			d := New(testObj)
			assert.Equalf(t, tt.want, d.CurrentBranchName(tt.args.path), "CurrentBranchName(%v)", tt.args.path)
		})
	}
}

func Test_defaultExec_Switch(t *testing.T) {
	type args struct {
		path   string
		branch string
	}
	tests := []struct {
		name     string
		args     args
		mockArgs outputArgs
		want     string
		wantErr  assert.ErrorAssertionFunc
	}{
		{"valid", args{"", ""}, outputArgs{"valid", true}, "valid", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testObj := new(MockedInterface)
			testObj.On("Output").Return(tt.mockArgs.text, tt.mockArgs.err)
			d := New(testObj)
			got, err := d.Switch(tt.args.path, tt.args.branch)
			if !tt.wantErr(t, err, fmt.Sprintf("Switch(%v, %v)", tt.args.path, tt.args.branch)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Switch(%v, %v)", tt.args.path, tt.args.branch)
		})
	}
}

func Test_defaultExec_Pull(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name     string
		args     args
		mockArgs outputArgs
		want     string
		wantErr  assert.ErrorAssertionFunc
	}{
		{"valid", args{""}, outputArgs{"valid", true}, "valid", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testObj := new(MockedInterface)
			testObj.On("Output").Return(tt.mockArgs.text, tt.mockArgs.err)
			d := New(testObj)
			got, err := d.Pull(tt.args.path)
			if !tt.wantErr(t, err, fmt.Sprintf("Pull(%v)", tt.args.path)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Pull(%v)", tt.args.path)
		})
	}
}

func Test_defaultExec_Clone(t *testing.T) {
	type args struct {
		repo string
		path string
	}
	tests := []struct {
		name     string
		args     args
		mockArgs outputArgs
		want     string
		wantErr  assert.ErrorAssertionFunc
	}{
		{"valid", args{"", ""}, outputArgs{"valid", true}, "valid", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testObj := new(MockedInterface)
			testObj.On("Output").Return(tt.mockArgs.text, tt.mockArgs.err)
			d := New(testObj)
			got, err := d.Clone(tt.args.repo, tt.args.path)
			if !tt.wantErr(t, err, fmt.Sprintf("Clone(%v, %v)", tt.args.repo, tt.args.path)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Clone(%v, %v)", tt.args.repo, tt.args.path)
		})
	}
}

func Test_defaultExec_Cleanup(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name     string
		args     args
		mockArgs outputArgs
		want     string
		wantErr  assert.ErrorAssertionFunc
	}{
		{"valid", args{""}, outputArgs{"valid", true}, "valid", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testObj := new(MockedInterface)
			testObj.On("Output").Return(tt.mockArgs.text, tt.mockArgs.err)
			d := New(testObj)
			got, err := d.Cleanup(tt.args.path)
			if !tt.wantErr(t, err, fmt.Sprintf("Cleanup(%v)", tt.args.path)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Cleanup(%v)", tt.args.path)
		})
	}
}

func Test_defaultExec_OriginUrl(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name     string
		args     args
		mockArgs outputArgs
		want     string
		wantErr  assert.ErrorAssertionFunc
	}{
		{"valid", args{""}, outputArgs{"valid", true}, "valid", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testObj := new(MockedInterface)
			testObj.On("Output").Return(tt.mockArgs.text, tt.mockArgs.err)
			d := New(testObj)
			got, err := d.OriginUrl(tt.args.path)
			if !tt.wantErr(t, err, fmt.Sprintf("OriginUrl(%v)", tt.args.path)) {
				return
			}
			assert.Equalf(t, tt.want, got, "OriginUrl(%v)", tt.args.path)
		})
	}
}
