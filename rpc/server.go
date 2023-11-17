package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

type Arith struct{}

type ArithRequest struct {
	A, B int
}
type ArithResponse struct {
	//乘积
	Pro int
	//商
	Quo int
	//余数
	Rem int
}

// 乘法
func (this *Arith) Multiply(req ArithRequest, resp *ArithResponse) error {
	resp.Pro = req.A * req.B
	return nil
}

// 商和余数
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	// 除
	res.Quo = req.A / req.B
	// 取模
	res.Rem = req.A % req.B
	return nil
}
func main() {
	rect := new(Arith)
	rpc.Register(rect)
	rpc.HandleHTTP()
	err := http.ListenAndServe("127.0.0.1:9999", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
