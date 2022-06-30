package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// 一段有问题的代码
func logicCode() {
	var c chan int // nil
	for {
		select {
		case v := <-c: // 阻塞
			fmt.Printf("recv from chan, value: %v\n", v)
		default:

		}
	}
}

// 在终端中，使用go tool pprof XXX.pprof报告
// 进入到pprof界面中，使用top -n查看详细函数的性能消耗

func main() {
	var isCPUpprof bool
	var isMempprof bool

	// 在终端执行 xx.exe -help 会出现"turn cpu pprof on"等提示
	flag.BoolVar(&isCPUpprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMempprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUpprof {
		f1, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err: %v", err)
			return
		}
		pprof.StartCPUProfile(f1)
		defer func() {
			pprof.StopCPUProfile()
			f1.Close()
		}()
	}

	for i := 0; i < 6; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)

	if isMempprof {
		f2, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(f2)
		f2.Close()
	}
}
