// Package service 业务逻辑层，具体的业务流程实现
package service

import (
	"strings"
	"sync"
	"time"

	"input_method/library"
	"input_method/model/module"
)

// Ime 输入法业务逻辑
type Ime struct {
	dictTrie  *module.Trie                   // 字典树
	dictWords map[string][]*library.DictWord // 所有词典
	mutex     sync.Mutex
}

// NewIme NewIme
func NewIme(args []string, batchSize int) *Ime {
	ime := &Ime{
		dictWords: make(map[string][]*library.DictWord),
	}

	// 若没输入词典文件，则打错误日志，输出错误信息
	if len(args) == 0 {
		library.LogService.Warning("No input dict file")
		return ime
	}

	// 1. 词典加载
	ime.dictMultiLoader(args, batchSize)

	// 2. 构建字典树
	ime.buildDictTrie()

	library.LogService.Notice("NewIme init done")
	return ime
}

// dictMultiLoader 词典加载并存储词典数据，采用并行加载，热更新
func (ime *Ime) dictMultiLoader(dictPathList []string, batchSize int) {
	sTime := time.Now()
	// 任务数量
	wg := &sync.WaitGroup{}
	wg.Add(len(dictPathList))
	// 限制并发数量
	ch := make(chan struct{}, batchSize)
	errCh := make([]error, len(dictPathList))
	for i, val := range dictPathList {
		ch <- struct{}{} // 占用缓冲区
		go func(index int, path string) {
			defer wg.Done()
			var dictWords library.DictData
			dictWords, errCh[index] = module.DictLoader(path)
			ime.lock()
			ime.dictWords[dictWords.Spell] = dictWords.Words
			ime.unLock()
			<-ch // 执行完缓冲区释放
		}(i, val)
	}
	wg.Wait()

	for i := range errCh {
		if errCh[i] != nil {
			library.LogService.Warning("DictLoader err: " + errCh[i].Error())
		}
	}
	library.LogService.Timecost("DictMultiLoader", sTime)
	library.LogService.Notice("DictMultiLoader done")
}

// buildDictTrie 将词典写入前缀树
func (ime *Ime) buildDictTrie() {
	sTime := time.Now()
	ime.dictTrie = module.Constructor()
	for spell, words := range ime.dictWords {
		ime.dictTrie.Insert(spell, words)
	}

	library.LogService.Timecost("BuildDictTrie", sTime)
	library.LogService.Notice("BuildDictTrie done")
}

// FindWords 根据拼音查找前缀树，并返回所有词
func (ime *Ime) FindWords(spell string) []string {
	var words []string
	searchWords := make([]*library.DictWord, 0)
	// 1. 首先判断 spell 是否符合纯拼音字符串（纯小写），如果不是则打印错误日志
	if len(spell) == 0 || !library.IsLowerAlphaStr(spell) {
		library.LogService.Warning("Spell error, please cheak your input")
		return words
	}
	library.LogService.Notice("Find words by spell: " + spell)

	sTime := time.Now()
	searchNode := ime.dictTrie.SearchPrefix(spell)
	library.LogService.Timecost("Find words Search", sTime)
	if searchNode == nil || !searchNode.GetIsEnd() {
		// 非精确匹配，直接从精确匹配返回的节点开始往下查找，不用从根节点重复查找
		sTime := time.Now()
		mergeChildren(searchNode, &searchWords)
		library.LogService.Timecost("Find words StartsWith", sTime)
		searchWords = ime.sort(searchWords)
		library.LogService.Notice("Find words: startsWith searchPrefix")
	} else {
		// 精确匹配
		searchWords = searchNode.GetWords()
		library.LogService.Notice("Find words: exact search")
	}

	// 只提取最终汉字
	words = make([]string, len(searchWords))
	for i, dictWord := range searchWords {
		if dictWord != nil {
			words[i] = dictWord.Word
		}
	}
	library.LogService.Timecost("Find words", sTime)
	library.LogService.Notice("Find words: " + strings.Join(words, ", "))
	return words
}

// mergeChildren 将子节点符合匹配规则的数据合并，此步骤自上而下加入数组，保证了字典序，但同层级节点无序
func mergeChildren(t *module.Trie, words *[]*library.DictWord) {
	if t == nil {
		return
	}
	if t.GetIsEnd() {
		// 必须是完整的拼音才行
		*words = append(*words, t.GetWords()...)
	}

	for _, child := range t.GetNodeList() {
		mergeChildren(child, words)
	}
}

// sort 候选词排序，由于最终只返回前10个，故采用自定义排序，实现了稳定的选择排序
func (ime *Ime) sort(srcWords []*library.DictWord) []*library.DictWord {
	if len(srcWords) == 0 {
		return []*library.DictWord{}
	}
	// 去重
	wordsMap := make(map[string]*library.DictWord)
	desWords := make([]*library.DictWord, 0, len(srcWords))
	for _, word := range srcWords {
		if _, ok := wordsMap[word.Word]; !ok {
			wordsMap[word.Word] = word
			desWords = append(desWords, word)
		}
	}
	srcWords = desWords
	// 返回最多 10 个
	length := 10
	if len(srcWords) < 10 {
		length = len(srcWords)
	}
	// 字母序即按照字符串大小比较即可，由于都是小写拼音，可直接比较
	for i := 0; i < length; i++ {
		largeIndex := i
		for j := i + 1; j < len(srcWords); j++ {
			// '[]' 优先级 大于 '*'
			// 字符串比较 "zhan" < "zhang" 为 TRUE，但字典序相反
			// 由于在字典树构建的时候，自上而下查询遵循了字典序原则；但同一层节点里是没有顺序的，因此还是需要判断字典序
			if srcWords[j].Frequency > srcWords[largeIndex].Frequency ||
				(srcWords[j].Frequency == srcWords[largeIndex].Frequency &&
					srcWords[j].Spell < srcWords[largeIndex].Spell) {
				largeIndex = j
			}
		}
		if i != largeIndex {
			swap(srcWords, i, largeIndex)
		}
	}
	return srcWords[:length]
}

// swap 实际上不是交换变量，而是将未排序数组后移
func swap(srcWords []*library.DictWord, index, largeIndex int) {
	tmp := srcWords[largeIndex]
	for i := largeIndex; i > index; i-- {
		srcWords[i] = srcWords[i-1]
	}
	srcWords[index] = tmp
}

// lock mutex 加锁
func (ime *Ime) lock() {
	ime.mutex.Lock()
}

// unLock mutex 释放锁
func (ime *Ime) unLock() {
	ime.mutex.Unlock()
}
