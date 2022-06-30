package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 文件日志结构体
type FileLogger struct {
	Level       LogLevel
	fileName    string
	filePath    string
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
}

// NewFileLogger 构造函数
func NewFileLogger(level, fname, fpath string, fsize int64) *FileLogger {
	loglevel, err := parseLogLevel(level)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       loglevel,
		fileName:    fname,
		filePath:    fpath,
		maxFileSize: fsize,
	}
	err = fl.initFile() // 按照文件路径和文件名打开文件
	if err != nil {
		panic(err)
	}
	return fl
}

// 根据指定的日志文件路径和文件名打开日志文件
func (f *FileLogger) initFile() error {
	fullFilePath := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFilePath+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	// 日志文件均已打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

// 判断日志文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err: %v\n", err)
		return false
	}
	// 如果当前文件大小 大于等于 规定的最大值，就返回true
	return fileInfo.Size() >= f.maxFileSize
}

// 记录日志的方法
func (f *FileLogger) log(lv, format string, arg ...interface{}) {
	n, _ := parseLogLevel(lv)
	msg := fmt.Sprintf(format, arg...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	if f.checkSize(f.fileObj) {
		newFile, err := f.splitFile(f.fileObj)
		if err != nil {
			return
		}
		f.fileObj = newFile
	}
	fmt.Fprintf(f.fileObj, "[%s][%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), lv, fileName, funcName, lineNo, msg)
	if n >= ERROR {
		if f.checkSize(f.errFileObj) {
			newFile, err := f.splitFile(f.errFileObj)
			if err != nil {
				return
			}
			f.errFileObj = newFile
		}
		fmt.Fprintf(f.errFileObj, "[%s][%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), lv, fileName, funcName, lineNo, msg)
	}
}

// 日志切割的方法
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割日志文件
	// 1. rename备份	xx.log --> xx.log.bak202208031709
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err: %v\n", err)
		return nil, err
	}
	nowStr := time.Now().Format("20060102150405")
	logName := path.Join(f.filePath, fileInfo.Name())      // 拿到当前的日志文件完整路径
	bakLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) // 拼接一个日志文件备份的名字
	os.Rename(logName, bakLogName)
	// 2. 关闭当前日志文件
	file.Close()
	// 3. 创建一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return nil, err
	}
	// 4. 将打开的新日志文件对象赋值给 f.fileObj
	return fileObj, nil
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

func (f *FileLogger) Debug(format string, arg ...interface{}) {
	if f.enable(DEBUG) {
		f.log("DEBUG", format, arg...)
	}
}

func (f *FileLogger) Info(format string, arg ...interface{}) {
	if f.enable(INFO) {
		f.log("INFO", format, arg...)
	}
}

func (f *FileLogger) Error(format string, arg ...interface{}) {
	if f.enable(ERROR) {
		f.log("ERROR", format, arg...)
	}
}

func (f *FileLogger) Warning(format string, arg ...interface{}) {
	if f.enable(WARNING) {
		f.log("WARNING", format, arg...)
	}
}

func (f *FileLogger) Fatal(format string, arg ...interface{}) {
	if f.enable(FATAL) {
		f.log("FATAL", format, arg...)
	}
}
