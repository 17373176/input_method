// Package library 文件包
package library

import (
	"path/filepath"
)

// FileExtensionFromPath 路径中获取文件后缀 extension
func FileExtensionFromPath(filePath string) string {
	return filepath.Ext(filePath)
}

// FileNameFromPath 路径中获取文件名，不包含后缀
func FileNameFromPath(filePath string) (name string) {
	_, fileName := filepath.Split(filePath)
	for i := len(fileName) - 1; i >= 0; i-- {
		if fileName[i] == '.' {
			return fileName[:i]
		}
	}
	return ""
}

// IsLowerAlphaStr 判断字符串是否全是小写字母
func IsLowerAlphaStr(str string) bool {
	for _, ch := range str {
		if ch < 'a' || ch > 'z' {
			return false
		}
	}
	return true
}
