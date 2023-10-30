package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

func main() {
	fmt.Print("请输入要扫描的IP地址: ")
	var ip string
	fmt.Scanln(&ip)
	fmt.Println("给你一点加载时间,你才知道是真的扫描而不是面向结果编程")
	var wg sync.WaitGroup
	ops := []int{}
	ports := make(chan int, 5000)
	for i := 1; i <= 5000; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done() //这句一定要放前面,确保被程序看到但是推迟了,如果放最后,return之后会直接看不见done
			addr := fmt.Sprintf("%s:%d", ip, j)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				return
			}
			ports <- j
			conn.Close()

		}(i)
	}
	fmt.Println("正在检测")
	wg.Wait()
	close(ports)
	for v := range ports {
		ops = append(ops, v)
	}
	sort.Ints(ops)
	for _, v := range ops {
		fmt.Printf("%d port is open\n", v)
	}
	fmt.Println("按下回车键以退出程序...")
	fmt.Scanln()
}
