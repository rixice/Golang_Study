package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// MysqlConfig MYSQL配置文件结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0. 参数校验
	// 0.1 传进来的data参数必须是指针ptr类型（因为需要在函数中对其进行赋值）
	tp := reflect.TypeOf(data)
	if tp.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer") // 格式化输出之后，返回一个error类型
		return
	}
	// 0.2 传进来的data参数必须是结构体类型指针（因为配置文件中各种键值对）
	if tp.Elem().Kind() != reflect.Struct { // 先取值（总不能int指针也通过）
		err = errors.New("data should be a struct pointer")
		return
	}
	// 1. 读文件得到字节类型byte数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//string(b) // 将文件内容转换成字符串
	line := strings.Split(string(b), "\n") // 将ReadFile()的文件内容按\n划分成string类型的切片
	fmt.Println(line)
	// 2. 一行一行地读数据
	var structName string
	for index, lineslice := range line {
		// 去掉字符串首尾的空格
		lineslice = strings.TrimSpace(lineslice)
		if len(lineslice) == 0 {
			continue // 如果是空行，就跳过
		}
		// 2.1 如果是注释，就忽略
		if strings.HasPrefix(lineslice, "#") || strings.HasPrefix(lineslice, ";") {
			continue
		}
		// 2.2 如果是[]，则代表是节（section）
		if strings.HasPrefix(lineslice, "[") {
			if lineslice[0] != '[' && lineslice[len(lineslice)-1] != ']' {
				err = fmt.Errorf("line%d syntax error", index+1)
				return
			}
			// 把这一行首尾的[]去掉，取中间的内容，并去除首尾空格
			sectionName := strings.TrimSpace(lineslice[1 : len(lineslice)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line%d syntax error", index+1)
				return
			}
			// 根据字符串的sectionName去data里面根据反射，找到对应的结构体
			for i := 0; i < tp.Elem().NumField(); i++ { // 因为是指针，所以要Elem()
				field := tp.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name
				}
			} // NumField()返回v中的字段数
		} else {
			// 2.3 如果都不是，就是以=分割的键值对
			// 1. 以等号分割这一行，等号左边是key，等号右边是value
			// 字符串内是否含有等号 或 是否非法以=开头
			if strings.Index(lineslice, "=") == -1 || strings.HasPrefix(lineslice, "=") {
				err = fmt.Errorf("line%v has syntax error", index+1)
				return
			}
			key_index := strings.Index(lineslice, "=") // 取'='的位置
			key := strings.TrimSpace(lineslice[:key_index])
			value := strings.TrimSpace(lineslice[key_index+1:])
			// 2. 根据structName去data里面把对应的嵌套结构体取出
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型信息

			if sType.Kind() != reflect.Struct { // 判断是否是嵌套结构体
				err = fmt.Errorf("data中的%v字段应该是一个结构体", structName)
				return
			}
			// 3. 遍历嵌套结构体的每一个字段，判断tag是不是等于key
			var fieldName string
			var field reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				field = sType.Field(i) // Tag信息是存储在Type信息中的
				if field.Tag.Get("ini") == key {
					// 找到对应的字段
					fieldName = field.Name
					break
				}
			}
			if len(fieldName) == 0 {
				// 在结构体中找不到对应的字符
				continue
			}
			// 4. 如果key = tag，给这个字段赋值
			// 4.1 根据fieldName去取出这个字段
			fileObj := sValue.FieldByName(fieldName)
			// 4.2 对结构体字段进行赋值
			fmt.Println(fieldName, field.Type.Kind())
			switch field.Type.Kind() {
			case reflect.String: // 当值是字符串时
				fileObj.SetString(value)
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
				var value_int int64
				value_int, err = strconv.ParseInt(value, 10, 64) // value默认是string，需要转换成int（10进制，64位）
				if err != nil {
					err = fmt.Errorf("line:%d value type error", index+1)
					return
				}
				fileObj.SetInt(value_int)
			case reflect.Bool: // 当值是布尔类型时
				var value_bool bool
				value_bool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", index+1)
					return
				}
				fileObj.SetBool(value_bool)
			}
		}
	}
	return
}

// ini配置文件解析器
func main() {
	var cfg config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err: %v\n", err)
		return
	}
	fmt.Println(cfg)
}
