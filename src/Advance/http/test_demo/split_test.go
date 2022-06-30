package split_string

import (
	"reflect"
	"testing"
)

// 单元测试
// func TestSplit(t *testing.T) {
// 	// 在这里写测试用例
// 	got := Split("babsbcbsdabf", "b")               // 实际拿到的结果
// 	want := []string{"", "a", "s", "c", "sda", "f"} // 期望拿到的结果
// 	if !reflect.DeepEqual(got, want) {
// 		// 测试用例失败了
// 		t.Errorf("want: %v\tbut got: %v\n", want, got)
// 	}
// 	// 成功，则输出PASS
// }

// 还可以写更多的测试用例
// func Test2Split(t *testing.T){}
///////////////////////////////////////////////////
// 测试组
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case 1": testCase{str: "abcasdsk", sep: "s", want: []string{"abca", "d", "k"}},
		"case 2": testCase{str: "Hello World", sep: "W", want: []string{"Hello ", "orld"}},
	}

	// testGroup := []testCase{
	// 	testCase{str: "呜呼哈呜呼", sep: "哈", want: []string{"呜呼", "呜呼"}},
	// 	testCase{str: "abcasdsk", sep: "s", want: []string{"abca", "d", "k"}},
	// 	testCase{str: "Hello World", sep: "W", want: []string{"Hello ", "orld"}},
	// }

	for name, tc := range testGroup {
		// got := Split(tc.str, tc.sep)
		// want := tc.want
		// if !reflect.DeepEqual(got, want) {
		// 	t.Errorf("want: %v\tbut got: %v\n", want, got)
		// }
		///////////////////////////////////////////////////
		// 子测试
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want: %#v\tbut got: %#v\n", tc.want, got)
			}
		})
	}
}

// BenchMark 基准测试
// 加上 -benchtime=? 可以修改测试时间，默认是一秒左右 <对于某些运行时间较长的程序使用>
func BenchmarkSplit(b *testing.B) {
	// 假设需要做一些耗时的无关操作
	// func XXX{}
	// b.ResetTimer()	重置计时器
	// b.SetParallelism	设置使用的CPU数
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
