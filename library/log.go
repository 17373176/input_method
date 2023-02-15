// Package library 公共包
package library

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Log 采用原生 log 包封装的日志包
type Log struct {
	logFile *os.File
}

// NewLog 初始化
func NewLog() *Log {
	logFile, err := os.OpenFile(LogDir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		LogService.Warning(err.Error())
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return &Log{
		logFile: logFile,
	}
}

// CloseLog 关闭日志文件
func (l *Log) CloseLog() {
	// 保证日志全部打印完毕
	time.Sleep(1 * time.Second)
	l.logFile.Close()
}

// Timecost 耗时日志
func (l *Log) Timecost(key string, sTime time.Time) {
	log.SetPrefix(TimeCostLog)
	log.Println(fmt.Sprint(key, TimeCostStr, time.Since(sTime).Milliseconds(), MillisecondsUnit))
}

// Notice 业务正常日志，用于追踪程序运行的日志信息
func (l *Log) Notice(logInfo string) {
	log.SetPrefix(NoticeLog)
	log.Println(logInfo)
}

// Warning 普通错误日志，不影响程序正常进行
func (l *Log) Warning(warning string) {
	log.SetPrefix(WarningLog)
	log.Println(warning)
}

/* // ------------------------------
// log/log4go 支持异步写入，使用 channel 达到加锁目的，线程安全
// Log 厂内 logit 封装
type Log struct {
	logger logit.Logger
	ctx    context.Context
}

// NewLog new Log
func NewLog() *Log {
	ctx := context.Background()
	webLogger, err := logit.NewLogger(ctx, logit.OptLogFileName(filepath.Join(LogDir, "ime.log")))
	// 不在控制台打印日志，日志级别为 info，意味着 info,warning,error都有
	if err != nil {
		panic(err.Error())
	}
	return &Log{
		logger: webLogger,
		ctx:    ctx,
	}
}

// Timecost 耗时日志
func (l *Log) Timecost(key string, sTime time.Time) {
	l.logger.Notice(l.ctx, fmt.Sprint(key, TimeCostStr, time.Since(sTime).Milliseconds(), MillisecondsUnit))
}

// Notice 业务正常日志，用于追踪程序运行的日志信息
func (l *Log) Notice(logInfo string) {
	l.logger.Notice(l.ctx, logInfo)
}

// Warning 普通错误日志，不影响程序正常进行
func (l *Log) Warning(warning string) {
	l.logger.Warning(l.ctx, warning)
}

// Fatal 遇到错误日志直接 return
func (l *Log) Fatal(warning string) {
	l.logger.Fatal(l.ctx, warning)
} */
