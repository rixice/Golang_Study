package main

import "fmt"

//import f "fmt"	// 包的别名
//import _ "fmt"	// 匿名导入包

// 包中的标识符（变量名\函数名\结构体\接口等）如果首字母是小写，则表示为private
// 首字母大写的标识符，则为public（e.g: fmt）
// 如果不想使用包内部的标识符（冗余），需要使用匿名导入

// 空接口
// interface：关键字
// interface{}: 空接口类型
func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16) // 避免动态申请内存
	m1["name"] = "张三"
	m1["age"] = "18"
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)

	show(false)
	show("hello world")
	show(nil)
}

// 空接口可以接收任意类型的值
func show(a interface{}) {
	_, ok := a.(string) // 类型断言
	if !ok {
		fmt.Println("该数据不是string类型")
	} else {
		fmt.Println("该数据是string类型")
	}
	fmt.Printf("%v\n", a)
}
