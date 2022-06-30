package main

import "time"

func main() {
	// log := NewConsoleLogger("WARNING")
	log := NewFileLogger("debug", "test", "./", 10*1024)
	i := 0
	for {
		log.Debug("[%d]这是一条DEBUG日志", i)
		log.Info("[%d]这是一条INFO日志", i)
		log.Warning("[%d]这是一条WARNING日志", i)
		log.Error("[%d]这是一条ERROR日志", i)
		log.Fatal("[%d]这是一条FATAL日志", i)
		time.Sleep(time.Second * 1)
		i++
	}
}
