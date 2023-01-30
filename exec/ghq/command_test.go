package ghq

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

func (m MockedInterface) FileExists(s string) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m MockedInterface) DirExists(s string) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m MockedInterface) Output(s string, s2 string, s3 ...string) (string, error) {
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

func TestValid(t *testing.T) {
	tests := []struct {
		name string
		mock bool
		want bool
	}{
		{"valid", true, true},
		{"invalid", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testObj := new(MockedInterface)
			testObj.On("NotError").Return(tt.mock)
			d := New(testObj)
			assert.Equalf(t, tt.want, d.Valid(), "Valid()")
			testObj.AssertExpectations(t)
		})
	}
}

func Test_defaultExec_List(t *testing.T) {
	tests := []struct {
		name     string
		mockArgs outputArgs
		want     []string
	}{
		{"valid", outputArgs{"a\nb", true}, []string{"a", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testObj := new(MockedInterface)
			testObj.On("Output").Return(tt.mockArgs.text, tt.mockArgs.err)
			d := New(testObj)
			assert.Equalf(t, tt.want, d.List(), "list()")
		})
	}
}

func Test_defaultExec_Get(t *testing.T) {
	type args struct {
		repo string
	}
	tests := []struct {
		name     string
		args     args
		mockArgs outputArgs
		want     string
		wantErr  assert.ErrorAssertionFunc
	}{
		{"valid", args{"repo"}, outputArgs{"main", true}, "main", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testObj := new(MockedInterface)
			testObj.On("Output").Return(tt.mockArgs.text, tt.mockArgs.err)
			d := New(testObj)
			got, err := d.Get(tt.args.repo)
			if !tt.wantErr(t, err, fmt.Sprintf("Get(%v)", tt.args.repo)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Get(%v)", tt.args.repo)
		})
	}
}

func Test_defaultExec_Path(t *testing.T) {
	type args struct {
		repo string
	}
	tests := []struct {
		name     string
		args     args
		mockArgs outputArgs
		want     string
		wantErr  assert.ErrorAssertionFunc
	}{
		{"valid", args{"repo"}, outputArgs{"main", true}, "main", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testObj := new(MockedInterface)
			testObj.On("Output").Return(tt.mockArgs.text, tt.mockArgs.err)
			d := New(testObj)
			got, err := d.Path(tt.args.repo)
			if !tt.wantErr(t, err, fmt.Sprintf("Path(%v)", tt.args.repo)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Path(%v)", tt.args.repo)
		})
	}
}
