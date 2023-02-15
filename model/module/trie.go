// Package module 前缀树子模块
package module

import (
	"fmt"
	"input_method/library"
)

// Trie 前缀树结构体
type Trie struct {
	isEnd    bool
	nodeList map[rune]*Trie      // 子节点数组，对于单词为小写英文的话，结构采用 map 映射，英文字母的 hash 值不会重复
	words    []*library.DictWord // 候选汉字数组
}

// Constructor Initialize
func Constructor() *Trie {
	return &Trie{
		nodeList: make(map[rune]*Trie),
	}
}

// Insert 递归迭代建树
// Inserts a word into the trie
func (t *Trie) Insert(spell string, words []*library.DictWord) {
	node := t
	for _, ch := range spell {
		ch -= 'a'
		if node.nodeList[ch] == nil {
			node.nodeList[ch] = &Trie{
				nodeList: make(map[rune]*Trie),
				words:    make([]*library.DictWord, 0, 5),
			}
		}
		node = node.nodeList[ch]
	}

	node.words = append(node.words, words...)
	node.isEnd = true
}

// SearchPrefix prefix 前缀匹配
func (t *Trie) searchPrefix(prefix string) *Trie {
	node := t
	if node == nil {
		fmt.Println("aa")
	}
	for _, ch := range prefix {
		ch -= 'a'
		if node.nodeList[ch] == nil {
			// 当前子节点为空
			return nil
		}
		node = node.nodeList[ch]
	}
	return node
}

// Search Search 精确查找
// Returns if the word is in the trie
func (t *Trie) Search(word string) ([]*library.DictWord, bool) {
	node := t.searchPrefix(word)
	if node != nil && node.isEnd {
		return node.words, true
	}
	return []*library.DictWord{}, false
}

// StartsWith StartsWith 前缀查找，返回所有匹配前缀的字
// Returns if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) []*library.DictWord {
	var words []*library.DictWord
	node := t.searchPrefix(prefix)
	if node == nil {
		return words
	}
	// 纯前缀匹配
	node.mergeChildren(&words)

	return words
}

// mergeChildren 将子节点符合匹配规则的数据合并，此步骤自上而下加入数组，保证了字典序，但同层级节点无序
func (t *Trie) mergeChildren(words *[]*library.DictWord) {
	if t.isEnd {
		// 必须是完整的拼音才行
		*words = append(*words, t.words...)
	}

	for _, child := range t.nodeList {
		child.mergeChildren(words)
	}
}
