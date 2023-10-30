package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	//创建读取对象
	reader := bufio.NewReader(conn)
	for {
		//建立存储客户端发送请求的数组
		rd := [512]byte{}
		//向数组写入内容
		n, err := reader.Read(rd[:])
		if err != nil {
			log.Fatalln(err)
		}
		//打印读取的内容
		fmt.Println(string(rd[:n]))
		//向客户端发送内容
		recv := fmt.Sprintf("你已成功发送,数据为:%s", string(rd[:n]))
		conn.Write([]byte(recv))
	}
}
func main() {
	listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		//建立连接
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go process(conn)
	}
}
