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

func TestNewLog(t *testing.T) {
	tests := []struct {
		name string
		want *Log
	}{
		{
			name: "new",
			want: NewLog(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLog(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_CloseLog(t *testing.T) {
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

func TestLog_Timecost(t *testing.T) {
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

func TestLog_Notice(t *testing.T) {
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

func TestLog_Warning(t *testing.T) {
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
