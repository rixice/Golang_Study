// 日志库作业
// 接口、文件操作
// 日志可以输出到终端，也可以输出到文件/kafka

// 需求：
// 1. 可以往不同的输出位置记录日志
// 2. 日志分为五种级别
// 3. 日志要支持开关控制(比如：开发时什么级别都输出，上线后只输出INFO及以下)
// 4. 完整的日志记录要包含有时间、行号、文件名、日志级别、日志信息
// 5. 日志文件要切割

package main

import "time"

// 测试日志库
func main() {
	log := NewConsoleLog("info")
	//log := NewFileLogger("INFO", "test.log", "./", 10*1024)
	i := 0
	for {
		log.Debug("%d_这是一条Debug日志", i)
		log.Error("这是一条Error日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Fatal("这是一条Fatal日志")
		i++
		time.Sleep(time.Second * 3)
	}

}
