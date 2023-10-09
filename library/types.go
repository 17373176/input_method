// Package library 包, types 类型
package library

// DictWord 单字结构
type DictWord struct {
	Spell     string
	Word      string
	Frequency int
}

// DictData 返回词典结构
type DictData struct {
	Spell string
	Words []*DictWord
}
