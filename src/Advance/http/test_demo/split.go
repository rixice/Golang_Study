package split_string

import (
	"strings"
)

// 切割字符串
// example:
// 传入abc，按b分割→[a c]

func Split(str string, sep string) []string {
	// str: "abcdebf"	sep:"b"
	// var ret []string
	// 使用切片这些时，一定要预估内存的使用量，减少内存的申请，优化性能
	var ret = make([]string, 0, strings.Count(str, sep)+1) // 减少了大量的内存申请
	index := strings.Index(str, sep)                       // 1
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+1:]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str) // 把切割完的最后一部分加入进来
	return ret
}
