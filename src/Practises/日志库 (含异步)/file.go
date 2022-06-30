package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger结构体
type FileLogger struct {
	Level       LogLevel
	fileName    string
	filePath    string
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
	logChan     chan *logMessage // 异步写日志
}

type logMessage struct {
	Level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	lineNo    int
}

// 切割日志文件
func (f *FileLogger) SplitFile(file *os.File) (*os.File, error) {
	// 1. rename备份文件
	fileInfo, _ := file.Stat()
	timeStr := time.Now().Format("20060102_150405")
	logName := path.Join(f.filePath, fileInfo.Name())
	bakName := fmt.Sprintf("%s.bak_%s", logName, timeStr)
	os.Rename(logName, bakName)
	// 2. 关闭当前文件
	file.Close()
	// 3. 打开新的文件
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return nil, err
	}
	return fileObj, nil
}

// 判断日志大小是否需要切割
func (f *FileLogger) CheckSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		return false
	}
	return fileInfo.Size() >= f.maxFileSize //大于最大值，则返回true要切割
}

// initFile 打开文件
func (f *FileLogger) initFile() error {
	fullFilePath := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open log_file failed, err: %v\n", err)
		return err
	}
	f.fileObj = fileObj
	errFileObj, err := os.OpenFile(fullFilePath+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log_file failed, err: %v\n", err)
		return err
	}
	f.errFileObj = errFileObj
	// 开启后台的goroutine去往文件里写日志
	for i := 0; i < 5; i++ {
		go f.writeLogBackground()
	}
	return nil
}

// NewFileLogger 构造函数
func NewFileLogger(level, fname, fpath string, fsize int64) *FileLogger {
	lv, err := parseLogLevel(level)
	if err != nil {
		panic(err)
	}
	file := &FileLogger{
		Level:       lv,
		fileName:    fname,
		filePath:    fpath,
		maxFileSize: fsize,
		logChan:     make(chan *logMessage, 50000),
	}
	err = file.initFile()
	if err != nil {
		panic(err)
	}
	return file
}

// 后台写日志
func (f *FileLogger) logBackground(levelStr, format string, arg ...interface{}) {
	level, _ := parseLogLevel(levelStr)
	msg := fmt.Sprintf(format, arg...)
	time := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(3)
	// 先把日志发送到通道中
	tmpLogMessage := &logMessage{
		Level:     level,
		msg:       msg,
		funcName:  funcName,
		fileName:  fileName,
		lineNo:    lineNo,
		timestamp: time,
	}
	select {
	case f.logChan <- tmpLogMessage:
	default:
		// 把日志丢弃，保证业务代码不出现阻塞
	}
}

func (f *FileLogger) writeLogBackground() {
	for {
		if f.CheckSize(f.fileObj) {
			newFileObj, err := f.SplitFile(f.fileObj)
			if err != nil {
				fmt.Printf("split file failed, err: %v\n", err)
			}
			f.fileObj = newFileObj
		}
		select {
		case logTmp := <-f.logChan:
			fmt.Fprintf(f.fileObj, "[%s][%v][%s|%s|%d] %s\n", logTmp.Level, logTmp.timestamp, logTmp.funcName, logTmp.fileName, logTmp.lineNo, logTmp.msg)
			if f.Level >= ERROR {
				if f.CheckSize(f.errFileObj) {
					newFileObj, err := f.SplitFile(f.fileObj)
					if err != nil {
						fmt.Printf("split file failed, err: %v\n", err)
					}
					f.errFileObj = newFileObj
				}
				fmt.Fprintf(f.errFileObj, "[%s][%s][%s|%s|%d] %s\n", logTmp.Level, logTmp.timestamp, logTmp.funcName, logTmp.fileName, logTmp.lineNo, logTmp.msg)
			}
		default:
			time.Sleep(time.Second * 1)
		}
	}
}

// 生成日志文件
// func (f *FileLogger) log(levelStr, format string, arg ...interface{}) {
// 	level, _ := parseLogLevel(levelStr)
// 	msg := fmt.Sprintf(format, arg...)
// 	time := time.Now().Format("2006-01-02 15:04:05")
// 	funcName, fileName, lineNo := getInfo(3)
// 	if f.CheckSize(f.fileObj) {
// 		newFileObj, err := f.SplitFile(f.fileObj)
// 		if err != nil {
// 			fmt.Printf("split file failed, err: %v\n", err)
// 		}
// 		f.fileObj = newFileObj
// 	}
// 	fmt.Fprintf(f.fileObj, "[%s][%v][%s|%s|%d] %s\n", levelStr, time, funcName, fileName, lineNo, msg)
// 	if level >= ERROR {
// 		if f.CheckSize(f.errFileObj) {
// 			newFileObj, err := f.SplitFile(f.fileObj)
// 			if err != nil {
// 				fmt.Printf("split file failed, err: %v\n", err)
// 			}
// 			f.errFileObj = newFileObj
// 		}
// 		fmt.Fprintf(f.errFileObj, "[%s][%s][%s|%s|%d] %s\n", levelStr, time, funcName, fileName, lineNo, msg)
// 	}
// }

// 判断日志等级
func (f *FileLogger) enable(lv LogLevel) bool {
	return f.Level <= lv
}

// DEBUG
func (f *FileLogger) Debug(format string, arg ...interface{}) {
	if f.enable(DEBUG) {
		f.logBackground("DEBUG", format, arg...)
	}
}

// INFO
func (f *FileLogger) Info(format string, arg ...interface{}) {
	if f.enable(INFO) {
		f.logBackground("INFO", format, arg...)
	}
}

// WARNING
func (f *FileLogger) Warning(format string, arg ...interface{}) {
	if f.enable(WARNING) {
		f.logBackground("WARNING", format, arg...)
	}
}

// ERROR
func (f *FileLogger) Error(format string, arg ...interface{}) {
	if f.enable(ERROR) {
		f.logBackground("ERROR", format, arg...)
	}
}

// FATAL
func (f *FileLogger) Fatal(format string, arg ...interface{}) {
	if f.enable(FATAL) {
		f.logBackground("FATAL", format, arg...)
	}
}
