package module

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"input_method/library"
)

func TestDictLoader(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   []*library.DictWord
		wantErr bool
	}{
		{
			name: "isn't lowerAlpha path",
			args: args{
				path: ".data/Zh.dat",
			},
			want:    "",
			want1:   nil,
			wantErr: true,
		},
		{
			name: "remote path",
			args: args{
				path: "http://xxx.baidu.com/zh.dat",
			},
			want:    "",
			want1:   nil,
			wantErr: true,
		},
		{
			name: "local path",
			args: args{path: "../../data/a.dat"},
			want: "a",
			want1: []*library.DictWord{
				{
					Word:      "啊",
					Frequency: 10,
					Spell:     "a",
				},
				{
					Word:      "阿",
					Frequency: 3,
					Spell:     "a",
				}},
			wantErr: false,
		},
		{
			name: "local file err",
			args: args{
				path: "zh.dat",
			},
			want:    "",
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := DictLoader(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("DictLoader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DictLoader() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DictLoader() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_localDictLoader(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []*library.DictWord
		wantErr bool
	}{
		{
			name: "valid path",
			args: args{path: "../../data/a.dat"},
			want: []*library.DictWord{
				{
					Word:      "啊",
					Frequency: 10,
				},
				{
					Word:      "阿",
					Frequency: 3,
				}},
			wantErr: false,
		},
		{
			name:    "invalid path",
			args:    args{path: "data/a.dat"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "parsing err",
			args:    args{path: "data/a.dat"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := localDictLoader(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("localDictLoader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("localDictLoader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpDictLoader(t *testing.T) {
	type args struct {
		path string
	}
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		_, err := io.WriteString(w, "展 3\n")
		if err != nil {
			fmt.Println(err)
		}
	}))
	defer server1.Close()
	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/plain")
		_, err := io.WriteString(w, "展 3")
		if err != nil {
			fmt.Println(err)
		}
	}))
	defer server2.Close()
	server3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")

	}))
	defer server3.Close()

	tests := []struct {
		name    string
		args    args
		want    []*library.DictWord
		wantErr bool
	}{
		{
			name: "valid path",
			args: args{
				path: server1.URL,
			},
			want: []*library.DictWord{
				{Word: "展", Frequency: 3},
			},
			wantErr: false,
		},
		{
			name: "404 StatusCode",
			args: args{
				path: server2.URL,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "parsing err",
			args: args{
				path: server3.URL,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "invalid path",
			args: args{
				path: "http://your.baidu.com",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := httpDictLoader(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("httpDictLoader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpDictLoader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkFilePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   bool
		wantErr bool
	}{
		{
			name: "valid path",
			args: args{
				path: "http://xxx.baidu.com/zh.dat",
			},
			want:    "zh",
			want1:   true,
			wantErr: false,
		},
		{
			name: "invalid path",
			args: args{
				path: ".data/zh.conf",
			},
			want:    "",
			want1:   false,
			wantErr: true,
		},
		{
			name: "isn't lowerAlpha path",
			args: args{
				path: ".data/Zh.dat",
			},
			want:    "",
			want1:   false,
			wantErr: true,
		},
		{
			name: "nil path",
			args: args{
				path: ".data/.dat",
			},
			want:    "",
			want1:   false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := checkFilePath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkFilePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkFilePath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkFilePath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_dictParsing(t *testing.T) {
	type args struct {
		file io.ReadCloser
	}
	file1, _ := os.Open("../../data/a.dat")
	defer file1.Close()
	file2, _ := os.Open("../../data/den.dat")
	defer file2.Close()
	file3, _ := os.Open("den.dat")
	defer file3.Close()
	tests := []struct {
		name       string
		args       args
		wantResult []*library.DictWord
	}{
		{
			name: "valid",
			args: args{
				file: file1,
			},
			wantResult: []*library.DictWord{
				{
					Word:      "啊",
					Frequency: 10,
				},
				{
					Word:      "阿",
					Frequency: 3,
				},
			},
		},
		{
			name: "invalid frequency",
			args: args{
				file: file2,
			},
			wantResult: nil,
		},
		{
			name: "invalid file",
			args: args{
				file: file3,
			},
			wantResult: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := dictParsing(tt.args.file)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("dictParsing() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
