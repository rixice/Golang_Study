package main

import (
	"fmt"
	"time"
)

func main() {
	f2()
}

func f1() {
	now := time.Now()
	nextYear, err := time.Parse("2006-01-02", "2019-08-04")
	if err != nil {
		fmt.Printf("parse time failed, err: %v\n", err)
		return
	}
	now = now.UTC()
	nextYear = nextYear.UTC()
	d := nextYear.Sub(now)
	fmt.Println(d)
}

// 时区
func f2() {
	now := time.Now() // 本地的时间
	fmt.Println(now)
	// 明天的这个时间
	// 按照指定格式取解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2019-07-01 14:14:33")
	// 按照东八区的时区和格式取解析一个字符串格式的时间
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed, err: %v\n", err)
		return
	}
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2019-07-01 14:14:33", loc)
	if err != nil {
		fmt.Printf("Parse time failed, err: %v\n", err)
		return
	}
	fmt.Println(timeObj)
	// 时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}
