package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":12345")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		//检测用户输入并传递给服务端(send)
		ip, _ := input.ReadString('\n')
		ipinfo := strings.Trim(ip, "\n\r")
		if strings.ToUpper(ipinfo) == "Q" {
			return
		}
		_, err := conn.Write([]byte(ipinfo)) //将用户的输入先强制转换成字节类型的切片,再发送给服务器
		if err != nil {
			log.Fatalln(err)
		}
		//接收服务端的消息(recv)
		buf := [512]byte{} //创建一个字节类型的数组([]里面写length就是数组,没写就是切片),用来存放服务端传递的数据(本例中服务端并未给客户端发送数据)
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(buf[:n]))
	}
}
