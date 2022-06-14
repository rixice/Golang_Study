package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与JSON（encode和decode）
// 1. 把Go语言中的结构体变量 --> json格式的字符串
// 2. json格式的字符串 --> Go语言中能够识别的结构体变量

type person struct {
	// 由于格式化的功能是在json包实现的，想要json包中拿到该变量，需要public
	Name string `json:"name" db:"name"` //表示在json的格式下，使用name这个字段名(需求要小写)<注意空格>
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "张三",
		Age:  18,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	// 反序列化
	str := `{"name":"李四","age":20}`
	var p2 person
	// 传指针：为了能在Unmarshal()内部去修改p2的值
	json.Unmarshal([]byte(str), &p2) // Unmarshal(data []byte, v any)
	fmt.Printf("%#v\n", p2)
}
