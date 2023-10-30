package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//创建路由对象,该函数的函数签名为  func mux.NewRouter() *mux.Router
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/websocket", myws)
	//所有访问本地2345端口的请求,都会被交由router这个路由器来处理
	if err := http.ListenAndServe("127.0.0.1:2345", router); err != nil {
		fmt.Println("Err:", err)
	}
}
