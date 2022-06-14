package main

import (
	"encoding/json"
	"fmt"
)

// 反射
// 想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，
// 必须传递变量地址才可以修改变量值
// 在反射中，要使用专有的Elem()来获取指针对应的值

type person struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
}

func main() {
	str := `{"name":"张三","age":18}`
	var p person
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)
}
