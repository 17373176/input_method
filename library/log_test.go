// Package library 公共包
package library

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

// TestNewLog 用于测试 NewLog 函数
func TestNewLog(t *testing.T) {
	LogService = NewLog("../log/", "ime.log")
	type args struct {
		logDir      string
		logFileName string
	}
	tests := []struct {
		name string
		args args
		want *Log
	}{
		{
			name: "new",
			want: LogService,
			args: args{
				logDir:      "../log/",
				logFileName: "ime.log"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogService; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestCloseLog 用于测试 NewLogFile 函数
func TestCloseLog(t *testing.T) {
	type fields struct {
		logFile *os.File
	}
	logFile, err := os.OpenFile(LogDir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "log",
			fields: fields{logFile: logFile},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				logFile: tt.fields.logFile,
			}
			l.CloseLog()
		})
	}
}

// TestTimecost 用于测试 Timecost 函数
func TestTimecost(t *testing.T) {
	type fields struct {
		logFile *os.File
	}
	type args struct {
		key   string
		sTime time.Time
	}
	logFile, err := os.OpenFile(LogDir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "log",
			fields: fields{logFile: logFile},
			args:   args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				logFile: tt.fields.logFile,
			}
			l.Timecost(tt.args.key, tt.args.sTime)
		})
	}
}

// TestNotice 用于测试 Notice 函数
func TestNotice(t *testing.T) {
	type fields struct {
		logFile *os.File
	}
	type args struct {
		logInfo string
	}
	logFile, err := os.OpenFile(LogDir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "log",
			fields: fields{logFile: logFile},
			args:   args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				logFile: tt.fields.logFile,
			}
			l.Notice(tt.args.logInfo)
		})
	}
}

// TestWarning 用于测试 Warning 函数
func TestWarning(t *testing.T) {
	type fields struct {
		logFile *os.File
	}
	type args struct {
		warning string
	}
	logFile, err := os.OpenFile(LogDir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "log",
			fields: fields{logFile: logFile},
			args:   args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				logFile: tt.fields.logFile,
			}
			l.Warning(tt.args.warning)
		})
	}
}
