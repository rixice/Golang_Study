package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net.http Client

func main() {
	data := url.Values{} // url encode
	urlObj, _ := url.Parse("http://127.0.0.1:8080/test/")
	data.Set("name", "张三")
	data.Set("age", "900")
	queryStr := data.Encode() // URL encode之后的URL
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, _ := http.NewRequest("GET", urlObj.String(), nil)
	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	// resp, err := http.Get("http://127.0.0.1:8080/test?name=hehe&&age=18")
	if err != nil {
		fmt.Println("Get url failed, err: ", err)
		return
	}
	/////////////////////////////////////////////////////////////////
	// 短连接，禁用KeepAlive的Client <需要定时刷新等>
	// tr := http.Transport{
	// 	DisableKeepAlives: true,
	// }
	// client := http.Client{
	// 	Transport: &tr,
	// }
	// resp, err = client.Do(req)
	/////////////////////////////////////////////////////////////////
	defer resp.Body.Close() // 一定要记得关闭resp.Body
	// 从resp中把服务端返回的数据读出来
	// var data []byte
	// resp.Body.Read(data)
	// resp.Body.Close()
	//////////////////////////////////
	b, err := ioutil.ReadAll(resp.Body) // 我在客户端读出服务端响应的body
	if err != nil {
		fmt.Println("read resp.body failed, err: ", err)
		return
	}
	fmt.Println(string(b))
}
