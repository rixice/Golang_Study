package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	var s1 = "haha"
	var s2 = s1 + "hehe"
	fmt.Println(s2)
	flag := false
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(j)
			if j == 3 {
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}

	/////////////////////////////////////////////

	// fallthrough 的用法
	n := 1
	switch n {
	case 1:
		fmt.Println("n=1")
		fallthrough
	case 2:
		fmt.Println("n=2")
	default:
		fmt.Println("ERROR")
	}

	////////////////////////////////////////////////

	// 初始化
	var arr_1 [8]bool
	fmt.Println(arr_1)

	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	// 多维数组
	var arr3 [2][3]int
	arr3 = [2][3]int{
		[3]int{1, 2, 3},
		[3]int{2, 3, 4},
	}
	fmt.Println(arr3)

	for _, v1 := range arr3 {
		for _, v2 := range v1 {
			fmt.Print(v2)
		}
	}

	///////////////////////////////////////////

	// 切片
	var ss1 []int
	var ss2 []string
	fmt.Println(s1, s2)

	ss1 = []int{1, 2, 3}
	ss2 = []string{"Hello", "World"}

	fmt.Println(ss1, ss2)
	fmt.Println(ss1 == nil)
	fmt.Println(ss2 == nil) // nil是null的意思

	// 长度和容量
	fmt.Printf("len: %d  cap: %d\n", len(ss2), cap(ss2))

	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9}
	s3 := a1[0:4] // 基于一个数组切割，左闭右开
	fmt.Println(s3)
	s4 := a1[1:]
	fmt.Println(s4)
	fmt.Printf("len: %d  cap: %d\n", len(s3), cap(s3))
	fmt.Printf("len: %d  cap: %d\n", len(s4), cap(s4))
	// 切片的容量：底层数组从切片的第一个元素到最后一个元素的数量
	// 切片的长度：切片元素的个数

	// 切片再切片
	s7 := s3[1:]
	fmt.Println(s7)
	fmt.Printf("len: %d  cap: %d\n", len(s7), cap(s7))

	// 切片是引用类型，都指向底层的一个数组
	// make()函数创造切片
	s10 := make([]int, 5, 10) // make([]T, len, cap)
	fmt.Printf("s10: %v  len: %d  cap: %d\n", s10, len(s10), cap(s10))
	// 切片不能直接比较，唯一合法的比较是和nil比较
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}

	////////////////////////////////////////////////
	// append() 为切片追加元素
	s11 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s11: %v  len: %d  cap: %d\n", s11, len(s11), cap(s11))
	s11 = append(s11, "广州") // 调用append函数，必须用原来的切片变量接收返回值
	fmt.Printf("s11: %v  len: %d  cap: %d\n", s11, len(s11), cap(s11))
	s11 = append(s11, "杭州", "珠海", "成都") // 到达cap上限后，cap翻倍
	fmt.Printf("s11: %v  len: %d  cap: %d\n", s11, len(s11), cap(s11))

	ss := []string{"武汉", "西安"}
	s11 = append(s11, ss...) // ... 表示解构
	fmt.Printf("s11: %v  len: %d  cap: %d\n", s11, len(s11), cap(s11))

	// copy
	s12 := make([]string, 9, 9)
	copy(s12, s11)
	fmt.Println(s12)

	// 将s12中的索引为3的元素删掉
	s13 := s12[:]
	s13 = append(s13[3:3], s13[4:]...) // 修改了底层数组
	fmt.Println(s13)
	/////////////////////////////////////////////////////////

	// 指针
	// 1. &:取地址
	nn := 10
	mm := &nn
	fmt.Println(mm)
	fmt.Printf("%T\n", mm)
	// 2. *:根据地址取值
	k := *mm
	fmt.Println(k)

	// make和new的区别
	// 1. make和new都是用来申请内存的
	// 2. new很少用，一般用来给 基本数据类型 申请内存，返回的是对应类型的指针
	// 3. make是用来给slice、map、chan申请内存的，返回的是对应的这三个类型本身

	// map
	var m1 map[string]int
	fmt.Println(m1 == nil) // 还没有初始化

	mm1 := make(map[string]int, 10) // 要估算好该map的容量，避免运行期间再动态扩容
	mm1["理想"] = 18
	mm1["籍无名"] = 35

	fmt.Println(mm1)
	delete(mm1, "理想")
	delete(mm1, "籍无名")
	fmt.Println(mm1["理想"])
	fmt.Println(mm1 == nil) // 已经开辟了内存，所以不为空

	v, ok := mm1["籍无名"]
	if !ok {
		fmt.Println("ERROR: 查无此key")
	} else {
		fmt.Println(v)
	}

	// 只遍历key
	for k, _ := range m1 {
		fmt.Println(k)
	}
	// 只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}

	////////////////////////////////////////

	// map和slice组合
	// 元素类型为map的切片
	var sm1 = make([]map[int]string, 10)
	// 没有对内部的map做初始化
	sm1[1] = make(map[int]string, 1)
	sm1[1][1] = "哈哈 "
	fmt.Println(sm1)

	var mi1 = make(map[string][]int, 10)
	mi1["北京"] = []int{10, 20, 30}
	fmt.Println(mi1)
	//////////////////////////////////////////////////////////////////

	// 作业：字符统计
	// 1. 判断字符串中汉字的数量
	// 难点：判断是否是汉字
	st1 := "Hello世界"
	var count int
	for _, v := range st1 {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	////////////////////////////////////////////////////////
	// 作业：how do you do 单词出现的次数
	// 2. 把字符串按照空格切割得到切片
	// 2.1 遍历切片存储到一个map
	// 2.2 累加出现的次数
	st3 := "how do you do"
	st4 := strings.Split(st3, " ")
	m3 := make(map[string]int, 10)
	for _, word := range st4 {

		if m3[word] == 0 {
			m3[word] = 1
		} else {
			m3[word]++
		}

	}
	fmt.Println(m3)
	fmt.Println(sum(10, 20))
	i, _ := f1()
	fmt.Println(i)
	i = f2(1, 2, "", "", true, false)
	fmt.Println(i)
}

// 返回值一定要定义类型
func sum(x int, y int) int {
	return x + y
}

// 多个返回值
func f1() (int, string) {
	return 1, "沙河"
}

// 参数类型简写(当参数中连续两个参数的类型一致时，我们可以将前面那个参数的类型省略)
func f2(x, y int, m, n string, i, j bool) int {
	return x + y
}

// 可变长参数
// 可变长参数必须放在函数参数的最后
// Go语言中函数没有默认参数这个概念
// Go语言中，函数传参传递的都是值（软连接）
func f3(x int, y ...int) { // y的类型是切片 []int
	fmt.Println(x)
	fmt.Println(y)
}
