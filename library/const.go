// Package library 公共常量
package library

const (
	// LogDir 日志目录
	LogDir = "./log/"
	// LogFileName 日志文件名
	LogFileName = "ime.log"
	// DictDir 词典目录
	DictDir = "./data"
)

// DictFileExt 词典文件后缀名
const DictFileExt = ".dat"

// URLRegular url 正则
const URLRegular = "https?://.+"

// BatchSize 最大并发数
const BatchSize = 8

// 日志相关
const (
	// NoticeLog notice 标记
	NoticeLog = "[NOTICE] "
	// WarningLog warning 标记
	WarningLog = "[WARNGING] "
	// FatalLog fatal 标记
	FatalLog = "[FATAL] "
	// TimeCostLog 耗时标记
	TimeCostLog = "[TIMECOST] "
	// TimeCostStr 耗时字符串
	TimeCostStr = " cost: "
	// MillisecondsUnit 毫秒单位
	MillisecondsUnit = "ms"
)
