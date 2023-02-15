// Package library 包, types 类型
package library

// Homonym 候选词
type Homonym struct {
	Word      string
	Frequency int
}

// DictWord 单字结构
type DictWord struct {
	Spell     string
	Word      string
	Frequency int
}
