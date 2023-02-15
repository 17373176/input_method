package module

import (
	"reflect"
	"testing"

	"library"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want *Trie
	}{
		{
			name: "new",
			want: Constructor(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Insert(t *testing.T) {
	type fields struct {
		isEnd    bool
		nodeList map[rune]*Trie
		words    []*library.DictWord
	}
	type args struct {
		spell string
		words []*library.DictWord
	}
	testWord := library.DictWord{Spell: "zhan", Word: "展", Frequency: 1}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "insert",
			fields: fields{
				isEnd:    false,
				nodeList: map[rune]*Trie{},
				words:    make([]*library.DictWord, 0),
			},
			args: args{
				spell: "zhan",
				words: []*library.DictWord{&testWord},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Trie{
				isEnd:    tt.fields.isEnd,
				nodeList: tt.fields.nodeList,
				words:    tt.fields.words,
			}
			tr.Insert(tt.args.spell, tt.args.words)
		})
	}
}

func TestTrie_searchPrefix(t *testing.T) {
	type fields struct {
		isEnd    bool
		nodeList map[rune]*Trie
		words    []*library.DictWord
	}
	type args struct {
		prefix string
	}
	testTrie := Constructor()
	testWord := library.DictWord{Spell: "zhan", Word: "展", Frequency: 1}
	testTrie.Insert("zhan", []*library.DictWord{&testWord})
	wantTrieExact := &Trie{isEnd: true, nodeList: make(map[rune]*Trie), words: []*library.DictWord{&testWord}}
	wantTriePrefix := &Trie{isEnd: false, nodeList: testTrie.nodeList['z'-'a'].nodeList, words: []*library.DictWord{}}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Trie
	}{
		{
			name: "searchPrefix exact",
			fields: fields{
				isEnd:    false,
				nodeList: testTrie.nodeList,
				words:    make([]*library.DictWord, 0),
			},
			args: args{
				prefix: "zhan",
			},
			want: wantTrieExact,
		},
		{
			name: "searchPrefix",
			fields: fields{
				isEnd:    false,
				nodeList: testTrie.nodeList,
				words:    make([]*library.DictWord, 0),
			},
			args: args{
				prefix: "z",
			},
			want: wantTriePrefix,
		},
		{
			name: "searchPrefix nil",
			fields: fields{
				isEnd:    false,
				nodeList: testTrie.nodeList,
				words:    make([]*library.DictWord, 0),
			},
			args: args{
				prefix: "zhang",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Trie{
				isEnd:    tt.fields.isEnd,
				nodeList: tt.fields.nodeList,
				words:    tt.fields.words,
			}
			if got := tr.searchPrefix(tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trie.searchPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Search(t *testing.T) {
	type fields struct {
		isEnd    bool
		nodeList map[rune]*Trie
		words    []*library.DictWord
	}
	type args struct {
		word string
	}
	testTrie := Constructor()
	testWord := library.DictWord{Spell: "zhan", Word: "展", Frequency: 1}
	testTrie.Insert("zhan", []*library.DictWord{&testWord})
	wantTrieExact := &Trie{isEnd: true, nodeList: make(map[rune]*Trie), words: []*library.DictWord{&testWord}}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*library.DictWord
		want1  bool
	}{
		{
			name: "search exact",
			fields: fields{
				isEnd:    false,
				nodeList: testTrie.nodeList,
				words:    make([]*library.DictWord, 0),
			},
			args: args{
				word: "zhan",
			},
			want:  wantTrieExact.words,
			want1: true,
		},
		{
			name: "search nil",
			fields: fields{
				isEnd:    false,
				nodeList: testTrie.nodeList,
				words:    make([]*library.DictWord, 0),
			},
			args: args{
				word: "z",
			},
			want:  []*library.DictWord{},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Trie{
				isEnd:    tt.fields.isEnd,
				nodeList: tt.fields.nodeList,
				words:    tt.fields.words,
			}
			got, got1 := tr.Search(tt.args.word)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trie.Search() = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Trie.Search() = %v, want1 %v", got1, tt.want1)
			}
		})
	}
}

func TestTrie_StartsWith(t *testing.T) {
	type fields struct {
		isEnd    bool
		nodeList map[rune]*Trie
		words    []*library.DictWord
	}
	type args struct {
		prefix string
	}
	testTrie := Constructor()
	testWord := library.DictWord{Spell: "zhan", Word: "展", Frequency: 1}
	testTrie.Insert("zhan", []*library.DictWord{&testWord})
	wantTrieExact := &Trie{isEnd: true, nodeList: make(map[rune]*Trie), words: []*library.DictWord{&testWord}}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*library.DictWord
	}{
		{
			name: "search exact",
			fields: fields{
				isEnd:    false,
				nodeList: testTrie.nodeList,
				words:    make([]*library.DictWord, 0),
			},
			args: args{
				prefix: "zhan",
			},
			want: wantTrieExact.words,
		},
		{
			name: "search nil",
			fields: fields{
				isEnd:    false,
				nodeList: testTrie.nodeList,
				words:    make([]*library.DictWord, 0),
			},
			args: args{
				prefix: "zhang",
			},
			want: nil,
		},
		{
			name: "search with",
			fields: fields{
				isEnd:    false,
				nodeList: testTrie.nodeList,
				words:    make([]*library.DictWord, 0),
			},
			args: args{
				prefix: "z",
			},
			want: []*library.DictWord{&testWord},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Trie{
				isEnd:    tt.fields.isEnd,
				nodeList: tt.fields.nodeList,
				words:    tt.fields.words,
			}
			if got := tr.StartsWith(tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trie.StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_mergeChildren(t *testing.T) {
	type fields struct {
		isEnd    bool
		nodeList map[rune]*Trie
		words    []*library.DictWord
	}
	type args struct {
		words *[]*library.DictWord
	}
	testTrie := Constructor()
	testWord := library.DictWord{Spell: "zhan", Word: "展", Frequency: 1}
	testTrie.Insert("zhan", []*library.DictWord{&testWord})
	words := []*library.DictWord{&testWord}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nil",
			fields: fields{
				isEnd:    false,
				nodeList: testTrie.nodeList,
				words:    make([]*library.DictWord, 0),
			},
			args: args{
				words: &words,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Trie{
				isEnd:    tt.fields.isEnd,
				nodeList: tt.fields.nodeList,
				words:    tt.fields.words,
			}
			tr.mergeChildren(tt.args.words)
		})
	}
}
