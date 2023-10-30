package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// 定义全局变量(package main的都能用)这里不能:= 只能var
var user_list = []string{}

// Upgrader是一个将http升级为websocket的函数
var wu = &websocket.Upgrader{
	/*
		指定了 WebSocket 连接在读取数据时的缓冲区大小，单位是字节。
		在 WebSocket 连接中，数据通常以消息的形式发送和接收，这个字段可以用于优化读取消息的性能。
		在这里，缓冲区大小被设置为 512 字节。
	*/
	ReadBufferSize: 512,
	/*
		指定了 WebSocket 连接在写入数据时的缓冲区大小，单位同样是字节。它用于优化将消息发送到 WebSocket 连接的性能。
		在这里，缓冲区大小也被设置为 512 字节。
	*/
	WriteBufferSize: 512,
	/*
		验证 WebSocket 连接的来源是否合法。WebSocket 连接可以受到同源策略的限制，而这个函数可以用来检查连接是否来自可信任的来源。
		在这个示例中，CheckOrigin 函数总是返回 true，表示接受来自任何来源的连接，
	*/
	CheckOrigin: func(r *http.Request) bool { return true },
}

type connection struct {
	//这里的websocket.Conn是一个结构体类型,整个ws是一个指向websocket.Conn的指针
	//ws代表了与客户端的 WebSocket 连接。通过这个字段，你可以执行与 WebSocket 连接相关的操作，如读取和写入消息，关闭连接等。
	ws *websocket.Conn
	//这里的sc是一个通道，用于传输客户端和服务器端的数据,能接收或者输出字节切片的通道,send channel
	sc   chan []byte
	data *Data
}

func myws(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{
		sc:   make(chan []byte, 256),
		ws:   ws,
		data: &Data{},
	}
	h.r <- c
	//reader的操作可能更频繁更消耗资源,所以放在主进程,而writer的操作不会占用太多资源,所以放在goroutine
	//这里的reader前面如果加了go程序会有问题
	go c.writer()
	c.reader()
	//这里的被defer的操作的逻辑有点问题,如果只是看这个程序的话那永远也不会执行注销操作,这里可以更改成看是否有error或者定时关闭
	defer func() {
		c.data.Type = "logout"
		user_list = del(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = c.data.User
		data_b, _ := json.Marshal(c.data)
		h.b <- data_b
		h.r <- c
	}()
}

// func del(slice []string, user string) []string {
// 	count := len(slice)
// 	if count == 0 {
// 		return slice
// 	}
// 	if count == 1 && slice[0] == user {
// 		return []string{}
// 	}
// 	n_slice := []string{}
// 	for i := range slice {
// 		if slice[i] == user && i == count {
// 			return []string{}
// 		} else if slice[i] == user {
// 			n_slice = append(n_slice, slice[i+1:]...)
// 			break
// 		}
// 	}
// 	fmt.Println(n_slice)
// 	return n_slice
// }

// 源del函数代码繁琐且错了,slice[i] == user && i == count ,这个永远不可能达到,这里的del函数是当有用户下线时,就从user_list里删除该用户
// 这里的方式是将创建新的数组n_slice,然后将user_list遍历一遍,不等于要被删除的用户就存进新的数组,最后返回新数组
func del(slice []string, user string) []string {
	var n_slice []string
	for _, val := range slice {
		if val != user {
			n_slice = append(n_slice, val)
		}
	}
	return n_slice
}

func (c *connection) writer() {
	for message := range c.sc {
		//websocket.TextMessage代表文本消息,是websocket包定义的一个常量
		/*
			TextMessage 表示文本数据消息。文本消息的负载被解释为 UTF-8 编码的文本数据
		*/
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}
func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			//文档给的h.r,但这里好像应该是h.u
			h.u <- c
			break
		}
		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			user_list = append(user_list, c.data.User)
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "user":
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "logout":
			c.data.Type = "logout"
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
			h.r <- c
		default:
			fmt.Println("=============deafult=================")
		}
	}
}
