package main

import (
	"fmt"
	"net/http"
)

func main() {
	//处理路由,当路由为/index时,触发myHandler的函数进行处理
	http.HandleFunc("/index", myHandler)
	//这里nil的意思是,除了127.0.0.1:1234/index以外的url访问,都会使用默认处理规则
	http.ListenAndServe("127.0.0.1:1234", nil)
}
func myHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//这里的w跟r是在有外界访问127.0.0.1:1234/index时默认传递的参数
	/*
		w http.ResponseWriter 是一个用于向客户端发送 HTTP 响应的接口。通过这个接口，您可以设置响应头、写入响应体等。
		在 myHandler 函数中，w 用于构建响应并发送给客户端。
		r *http.Request 是一个表示 HTTP 请求的结构体。它包含了关于客户端请求的信息，如请求方法、URL、请求头、请求体等。
		在 myHandler 函数中，r 用于访问客户端的请求信息，以便根据请求执行相应的操作。
	*/
	//简单来说w是服务端对客户端进行的操作,而r是客户端发来的请求对象,以下是对r的操作
	fmt.Println(r.RemoteAddr, ":连接成功")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	//这里的w是服务端对客户端的操作
	w.Write([]byte("这是服务器发来的消息"))
}
