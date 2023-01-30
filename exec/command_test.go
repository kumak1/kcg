package exec

import (
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
			if got := FileExists(tt.args.path); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
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
			if got := DirExists(tt.args.path); got != tt.want {
				t.Errorf("DirExists() = %v, want %v", got, tt.want)
			}
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
		wantErr bool
	}{
		{"valid", args{"", "echo", []string{"a"}}, "a", false},
		{"invalid", args{"", "invalid", []string{}}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Output(tt.args.path, tt.args.name, tt.args.arg...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Output() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Output() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotError(t *testing.T) {
	type args struct {
		path string
		name string
		arg  []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{"", "ls", []string{}}, true},
		{"invalid", args{"", "invalid", []string{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotError(tt.args.path, tt.args.name, tt.args.arg...); got != tt.want {
				t.Errorf("NotError() = %v, want %v", got, tt.want)
			}
		})
	}
}
