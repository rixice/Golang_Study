package main

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint64

const (
	UNKNOWN LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

type ConsoleLogger struct {
	level LogLevel
}

// Level字符串转LogLevel类型
func parseLogLevel(str string) (LogLevel, error) {
	str = strings.ToLower(str)
	switch str {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		return UNKNOWN, nil
	}
}

// 构造函数
func NewConsoleLogger(levelStr string) *ConsoleLogger {
	Level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &ConsoleLogger{
		level: Level,
	}
}

// 生成日志(终端)
func (c *ConsoleLogger) log(level, format string, arg ...interface{}) {
	msg := fmt.Sprintf(format, arg...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(2)
	fmt.Printf("[%s][%s][%s|%s|%d] %s\n", now.Format("2006-01-02 15:04:05"), level, funcName, fileName, lineNo, msg)
}

// 检查日志等级是否输出
func (c *ConsoleLogger) enable(lv LogLevel) bool {
	return c.level <= lv
}

// DEBUG
func (c *ConsoleLogger) Debug(format string, arg ...interface{}) {
	if c.enable(DEBUG) {
		c.log("DEBUG", format, arg...)
	}
}

// INFO
func (c *ConsoleLogger) Info(format string, arg ...interface{}) {
	if c.enable(INFO) {
		c.log("INFO", format, arg...)
	}
}

// WARNING
func (c *ConsoleLogger) Warning(format string, arg ...interface{}) {
	if c.enable(WARNING) {
		c.log("WARNING", format, arg...)
	}
}

// ERROR
func (c *ConsoleLogger) Error(format string, arg ...interface{}) {
	if c.enable(ERROR) {
		c.log("ERROR", format, arg...)
	}
}

func (c *ConsoleLogger) Fatal(format string, arg ...interface{}) {
	if c.enable(FATAL) {
		c.log("FATAL", format, arg...)
	}
}

// 获取日志文件信息
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed, please check !")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return funcName, fileName, lineNo
}
