package main

import "encoding/json"

var h = hub{
	c: make(map[*connection]bool),
	b: make(chan []byte),
	r: make(chan *connection),
	u: make(chan *connection),
}

type hub struct {
	//condition 状态
	c map[*connection]bool
	//broadcast 广播消息
	b chan []byte
	//register 注册
	r chan *connection
	//unregister 注销
	u chan *connection
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.r:
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			c.sc <- data_b
		case c := <-h.u:
			//这里的源代码逻辑有问题
			delete(h.c, c)
			//close(c.sc)

		case data := <-h.b:
			//这里的range遍历,c获得的是*connection这个键
			//c: make(map[*connection]bool)    range遍历映射时会获取键而不是键值对
			for c := range h.c {
				select {
				//先将data(这里的data是[]byte)传递给c.sc这个channel,如果传递不了则执行default的操作
				case c.sc <- data:
				default:
					delete(h.c, c)
					close(c.sc)
				}
			}
		}

	}
}
