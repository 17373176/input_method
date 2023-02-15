// Package controller 控制层，决定整体流程
package controller

import (
	"input_method/library"
	"input_method/model/service"
)

// InputMethod 输入法结构
type InputMethod struct {
	ime *service.Ime
}

// NewInputMethod 创建一个 InputMethod 实例
// 如果词典文件格式有误，忽略格式有误的文件
func NewInputMethod(args []string) *InputMethod {
	// 加载词典并构造字典树
	return &InputMethod{
		ime: service.NewIme(args, library.BatchSize),
	}
}

// FindWords 根据输入的拼音返回对应的汉字
func (im *InputMethod) FindWords(spell string) (words []string) {
	// Your code here
	return im.ime.FindWords(spell)
}
