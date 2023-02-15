// Package library 文件包
package library

import "testing"

func TestFileExtensionFromPath(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid",
			args: args{filePath: "./data/a.dat"},
			want: ".dat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExtensionFromPath(tt.args.filePath); got != tt.want {
				t.Errorf("FileExtensionFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileNameFromPath(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name     string
		args     args
		wantName string
	}{
		{
			name:     "valid",
			args:     args{filePath: "./data/a.dat"},
			wantName: "a",
		},
		{
			name:     "invalid",
			args:     args{filePath: "./data"},
			wantName: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotName := FileNameFromPath(tt.args.filePath); gotName != tt.wantName {
				t.Errorf("FileNameFromPath() = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}

func TestIsLowerAlphaStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "str is 'Zh'",
			args: args{"Zh"},
			want: false,
		},
		{
			name: "str is 'zh'",
			args: args{"zh"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLowerAlphaStr(tt.args.str); got != tt.want {
				t.Errorf("IsLowerAlphaStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
