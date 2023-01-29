package kcg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	assert.True(t, validGroup("valid", []string{"valid"}))
	assert.False(t, validGroup("invalid", []string{"valid"}))
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
