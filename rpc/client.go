package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type ArithRequest struct {
	A, B int
}

// 返回给客户端的结果
type ArithResponse struct {
	// 乘积
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}

func main() {
	conn, _ := rpc.DialHTTP("tcp", "127.0.0.1:9999")
	req := ArithRequest{9, 2}
	var res ArithResponse
	err2 := conn.Call("Arith.Multiply", req, &res)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)
	err3 := conn.Call("Arith.Divide", req, &res)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Printf("%d / %d 商 %d，余数 = %d\n", req.A, req.B, res.Quo, res.Rem)
}
