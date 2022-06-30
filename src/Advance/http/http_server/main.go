package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net.http server端
// HTTP：规定了浏览器和网站服务器之间通信的规则

func f1(response http.ResponseWriter, handler *http.Request) {
	// str := "<h1>hello world</h1>"
	b, _ := ioutil.ReadFile("./test.html")
	// response.Write([]byte(str))
	response.Write(b)
}

func f2(response http.ResponseWriter, handler *http.Request) {
	// 对于GET请求，参数都放在URL上，请求体是没有数据的
	// fmt.Println(handler.URL)
	queryParam := handler.URL.Query() // 自动帮我们识别URL中的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(handler.Method)
	fmt.Println(ioutil.ReadAll(handler.Body)) // 在服务端打印客户端发来的Body
	response.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/test/", f2)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
