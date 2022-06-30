package main

import (
	"flag"
	"fmt"
)

// 定义命令行flag参数
// 有以下两种常用的定义命令行flag参数的方法

// flag.Type()
// 基本格式如下：
// flag.Type(flag名，默认值，help消息)*Type
// 例如我们要定义姓名、年龄
// name := flag.String("name", "张三", "姓名")
// age := flag.Int("age", 18, "年龄")

func main() {
	// fmt.Println(os.Args)
	// fmt.Printf("%T\n", os.Args)	--> []string
	//////////////////////////////////////////////
	// 创建一个标志位参数
	// name := flag.String("name", "张三", "请输入姓名") // 返回的是地址
	// cTime := flag.Duration("time", time.Second, "时间") // 时间标志位<h,m,s等>

	// flag.TypeVar()
	var name string
	flag.StringVar(&name, "name", "张三", "请输入名字")

	// 使用flag
	flag.Parse() // 一定要有，无法使用标志位，保持默认输出 <并且要先解析，后面才能使用变量>
	// fmt.Println(name)

	// e.g: 输入 xx.exe a，结果为：[a]  1  0
	fmt.Println(flag.Args())  // 返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  // 返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) // 返回使用的命令行参数个数
	// fmt.Println(*cTime)
}
