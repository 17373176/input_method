// Package controller 控制层，决定整体流程
package controller

import (
	"reflect"
	"testing"

	"input_method/model/service"
)

func TestNewInputMethod(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want *InputMethod
	}{
		{
			name: "new",
			args: args{[]string{"../data/a.dat"}},
			want: NewInputMethod([]string{"../data/a.dat"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputMethod(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInputMethod_FindWords(t *testing.T) {
	type fields struct {
		ime *service.Ime
	}
	type args struct {
		spell string
	}
	testIme := service.NewIme([]string{"../data/a.dat"}, 2)
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantWords []string
	}{
		{
			name:      "find",
			fields:    fields{ime: testIme},
			args:      args{spell: "a"},
			wantWords: []string{"啊", "阿"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			im := &InputMethod{
				ime: tt.fields.ime,
			}
			if gotWords := im.FindWords(tt.args.spell); !reflect.DeepEqual(gotWords, tt.wantWords) {
				t.Errorf("InputMethod.FindWords() = %v, want %v", gotWords, tt.wantWords)
			}
		})
	}
}
