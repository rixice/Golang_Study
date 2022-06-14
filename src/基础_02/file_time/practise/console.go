package main

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

// 往终端写日志相关内容
type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	// 定义日志级别 0~5
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s) // 强制转换小写
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "error":
		return ERROR, nil
	case "warning":
		return WARNING, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

type ConsoleLogger struct {
	Level LogLevel
}

// NewConsoleLog 构造函数
func NewConsoleLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

// 记录日志的方法
func (c ConsoleLogger) log(level string, format string, arg ...interface{}) { // 模仿fmt.printf的格式，使用户可以格式化字符串
	msg := fmt.Sprintf(format, arg...) // Sprintf返回一个字符串，不输出
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s][%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), level, fileName, funcName, lineNo, msg)
}

// 判断是否记录的日志等级
func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return c.Level <= logLevel
}

func (c ConsoleLogger) Debug(format string, arg ...interface{}) {
	if c.enable(DEBUG) {
		c.log("DEBUG", format, arg...)
	}
}

func (c ConsoleLogger) Info(format string, arg ...interface{}) {
	if c.enable(INFO) {
		c.log("INFO", format, arg...)
	}
}

func (c ConsoleLogger) Error(format string, arg ...interface{}) {
	if c.enable(ERROR) {
		c.log("ERROR", format, arg...)
	}
}

func (c ConsoleLogger) Warning(format string, arg ...interface{}) {
	if c.enable(WARNING) {
		c.log("WARNING", format, arg...)
	}
}

func (c ConsoleLogger) Fatal(format string, arg ...interface{}) {
	if c.enable(FATAL) {
		c.log("FATAL", format, arg...)
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed!")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
