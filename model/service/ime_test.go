package service

import (
	"reflect"
	"testing"

	"input_method/library"
	"input_method/model/module"
)

func TestNewIme(t *testing.T) {
	type args struct {
		args      []string
		batchSize int
	}
	testIme := NewIme([]string{"../../data/a.dat"}, 2)
	tests := []struct {
		name string
		args args
		want *Ime
	}{
		{
			name: "empty args error",
			args: args{
				args:      nil,
				batchSize: 2,
			},
			want: &Ime{
				dictWords: make(map[string][]*library.DictWord),
			},
		},
		{
			name: "valid",
			args: args{
				args:      []string{"../../data/a.dat"},
				batchSize: 2,
			},
			want: testIme,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIme(tt.args.args, tt.args.batchSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIme_dictMultiLoader(t *testing.T) {
	type fields struct {
		dictTrie  *module.Trie
		dictWords map[string][]*library.DictWord
	}
	type args struct {
		dictPathList []string
		batchSize    int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "dict dict words, batch size > 1",
			fields: fields{
				dictTrie:  module.Constructor(),
				dictWords: make(map[string][]*library.DictWord),
			},
			args: args{
				dictPathList: []string{"../../data/zhan.dat"},
				batchSize:    2,
			},
		},
		{
			name: "dict err",
			fields: fields{
				dictTrie:  module.Constructor(),
				dictWords: make(map[string][]*library.DictWord),
			},
			args: args{
				dictPathList: []string{"zhan.dat"},
				batchSize:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ime := &Ime{
				dictTrie:  tt.fields.dictTrie,
				dictWords: tt.fields.dictWords,
			}
			ime.dictMultiLoader(tt.args.dictPathList, tt.args.batchSize)
		})
	}
}

func TestIme_buildDictTrie(t *testing.T) {
	type fields struct {
		dictTrie  *module.Trie
		dictWords map[string][]*library.DictWord
	}
	testWord := library.DictWord{Spell: "z", Word: "长", Frequency: 1}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "word build",
			fields: fields{
				dictTrie:  module.Constructor(),
				dictWords: map[string][]*library.DictWord{"z": {&testWord}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ime := &Ime{
				dictTrie:  tt.fields.dictTrie,
				dictWords: tt.fields.dictWords,
			}
			ime.buildDictTrie()
		})
	}
}

func TestIme_FindWords(t *testing.T) {
	type fields struct {
		dictTrie  *module.Trie
		dictWords map[string][]*library.DictWord
	}
	testTrie := module.Constructor()
	testWord := library.DictWord{Spell: "z", Word: "长", Frequency: 1}
	testTrie.Insert("z", []*library.DictWord{&testWord})
	type args struct {
		spell string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "empty spell",
			fields: fields{
				dictTrie:  module.Constructor(),
				dictWords: make(map[string][]*library.DictWord),
			},
			args: args{
				spell: "",
			},
		},
		{
			name: "error spell",
			fields: fields{
				dictTrie:  module.Constructor(),
				dictWords: make(map[string][]*library.DictWord),
			},
			args: args{
				spell: "LOVE",
			},
		},
		{
			name: "spell not search",
			fields: fields{
				dictTrie: module.Constructor(),
				dictWords: map[string][]*library.DictWord{
					"zhan": {0: &testWord},
				},
			},
			args: args{
				spell: "zh",
			},
			want: []string{},
		},
		{
			name: "spell exact search",
			fields: fields{
				dictTrie: testTrie,
				dictWords: map[string][]*library.DictWord{
					"z": {0: &library.DictWord{Spell: "z", Word: "长", Frequency: 1}},
				},
			},
			args: args{
				spell: "z",
			},
			want: []string{"长"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ime := &Ime{
				dictTrie:  tt.fields.dictTrie,
				dictWords: tt.fields.dictWords,
			}
			if got := ime.FindWords(tt.args.spell); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ime.FindWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIme_sort(t *testing.T) {
	type fields struct {
		dictTrie  *module.Trie
		dictWords map[string][]*library.DictWord
	}
	type args struct {
		srcWords []*library.DictWord
	}
	test1 := library.DictWord{Spell: "z", Word: "长", Frequency: 1}
	test2 := library.DictWord{Spell: "zh", Word: "张", Frequency: 2}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*library.DictWord
	}{
		{
			name: "empty",
			fields: fields{
				dictTrie:  module.Constructor(),
				dictWords: make(map[string][]*library.DictWord),
			},
			args: args{
				srcWords: make([]*library.DictWord, 0),
			},
			want: []*library.DictWord{},
		},
		{
			name: "len(srcWords) < 10",
			fields: fields{
				dictTrie:  module.Constructor(),
				dictWords: make(map[string][]*library.DictWord),
			},
			args: args{
				srcWords: []*library.DictWord{
					&test1, &test2,
				},
			},
			want: []*library.DictWord{
				&test2, &test1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ime := &Ime{
				dictTrie:  tt.fields.dictTrie,
				dictWords: tt.fields.dictWords,
			}
			if got := ime.sort(tt.args.srcWords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ime.sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swap(t *testing.T) {
	type args struct {
		srcWords   []*library.DictWord
		index      int
		largeIndex int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "swap case",
			args: args{
				srcWords: []*library.DictWord{
					{Word: "你", Spell: "ni", Frequency: 1},
					{Word: "好", Spell: "hao", Frequency: 1},
				},
				index:      0,
				largeIndex: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.args.srcWords, tt.args.index, tt.args.largeIndex)
		})
	}
}
