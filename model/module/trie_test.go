package module

import (
	"reflect"
	"testing"

	"input_method/library"
)

// TestConstructor 测试构造函数
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

// TestGetIsEnd 测试是否为结束
func TestGetIsEnd(t *testing.T) {
	type fields struct {
		isEnd    bool
		nodeList map[rune]*Trie
		words    []*library.DictWord
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "not nil",
			fields: fields{},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Trie{
				isEnd:    tt.fields.isEnd,
				nodeList: tt.fields.nodeList,
				words:    tt.fields.words,
			}
			if got := tr.GetIsEnd(); got != tt.want {
				t.Errorf("Trie.GetIsEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetNodeList 测试获取节点列表
func TestGetNodeList(t *testing.T) {
	type fields struct {
		isEnd    bool
		nodeList map[rune]*Trie
		words    []*library.DictWord
	}
	node := make(map[rune]*Trie)
	tests := []struct {
		name   string
		fields fields
		want   map[rune]*Trie
	}{
		{
			name: "nil",
			fields: fields{
				nodeList: node,
			},
			want: node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Trie{
				isEnd:    tt.fields.isEnd,
				nodeList: tt.fields.nodeList,
				words:    tt.fields.words,
			}
			if got := tr.GetNodeList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trie.GetNodeList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetWords 测试获取词表
func TestGetWords(t *testing.T) {
	type fields struct {
		isEnd    bool
		nodeList map[rune]*Trie
		words    []*library.DictWord
	}
	testWord := library.DictWord{Spell: "z", Word: "展", Frequency: 1}
	tests := []struct {
		name   string
		fields fields
		want   []*library.DictWord
	}{
		{
			name: "not nil",
			fields: fields{
				isEnd: true,
				nodeList: map[rune]*Trie{
					'z': {
						isEnd: true,
						words: []*library.DictWord{&testWord},
					},
				},
				words: []*library.DictWord{&testWord},
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
			if got := tr.GetWords(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trie.GetWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestInsert 测试插入节点
func TestInsert(t *testing.T) {
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

// TestSearchPrefix 测试前缀查询
func TestSearchPrefix(t *testing.T) {
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
			if got := tr.SearchPrefix(tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trie.searchPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
