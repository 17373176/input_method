// Package module 词典加载模块
package module

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"input_method/library"
)

// DictLoader 单个加载
func DictLoader(path string) (library.DictData, error) {
	var dictData library.DictData
	var dictWords []*library.DictWord
	// 判断文件路径是否包含http:// 或 https:// 前缀
	spell, isRemote, err := checkFilePath(path)
	if err != nil {
		return dictData, fmt.Errorf("checkFilePath err: %w", err)
	}

	dictWords, err = loadFuncMap[isRemote](path)
	if err != nil {
		return dictData, fmt.Errorf("dictLoader err: %w", err)
	}
	// spell 添加
	for i := range dictWords {
		dictWords[i].Spell = spell
	}
	// 先按照频次排序，需要用到稳定排序，否则位置会变更
	sort.SliceStable(dictWords, func(i, j int) bool {
		return dictWords[i].Frequency > dictWords[j].Frequency
	})
	dictData = library.DictData{
		Spell: spell,
		Words: dictWords,
	}
	return dictData, nil
}

var loadFuncMap = map[bool]func(string) ([]*library.DictWord, error){
	true:  httpDictLoader,
	false: localDictLoader,
}

// localDictLoader 从本地加载
func localDictLoader(path string) ([]*library.DictWord, error) {
	file, err := os.Open(path)
	defer file.Close()
	var result []*library.DictWord
	if err != nil {
		return result, fmt.Errorf("open file: %s, err: %w", path, err)
	}
	result = dictParsing(file)
	return result, nil
}

// httpDictLoader 从远程加载
func httpDictLoader(path string) ([]*library.DictWord, error) {
	var result []*library.DictWord

	response, err := http.Get(path)
	if err != nil {
		return result, fmt.Errorf("request: %s, err: %w", path, err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return result, fmt.Errorf("request: %s, err: status code is not 200", path)
	}

	result = dictParsing(response.Body)
	return result, nil
}

// checkFilePath 文件路径检查，返回文件名即拼音
func checkFilePath(path string) (string, bool, error) {
	// 正则匹配 http 和 https 路径
	match := library.RegexMatch.MatchString(path)

	// 获取文件名
	fileName := library.FileNameFromPath(path)
	if len(fileName) == 0 {
		return "", match, errors.New("fileName Error")
	}

	// 如果不是全小写字母，则报错
	if !library.IsLowerAlphaStr(fileName) {
		return "", match, errors.New("fileName isn't lowerAlphaStr")
	}

	// 文件后缀必须为.dat
	if library.FileExtensionFromPath(path) != library.DictFileExt {
		return "", match, errors.New("fileExt Error")
	}
	return fileName, match, nil
}

// dictParsing 词典文件解析，若文件格式有误，忽略该文件
func dictParsing(file io.ReadCloser) (result []*library.DictWord) {
	reader := bufio.NewReader(file)
	for {
		readLine, err := reader.ReadString('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			return result
		}
		dataLine := strings.TrimRight(readLine, "\n")
		kv := strings.Split(dataLine, " ")
		if len(kv) == 2 {
			word := kv[0]
			// 验证单词频次是否是int，不是则忽略该字
			frequency, err := strconv.Atoi(kv[1])
			if err != nil {
				continue
			}
			w := &library.DictWord{
				Word:      word,
				Frequency: frequency,
			}
			result = append(result, w)
		}
		// 为了保证最后一行能被读到
		if errors.Is(err, io.EOF) {
			break
		}
	}
	return result
}
