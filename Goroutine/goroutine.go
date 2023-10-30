package main

import (
	"fmt"
	"math/rand"
)

type Result struct {
	job *Job
	sum int
}
type Job struct {
	Id            int
	Random_number int
}

func main() {
	//先建立需要的channel,分别是将job传给工人的channel和结果的channel
	send_channel := make(chan *Job, 128)
	result_channel := make(chan *Result, 128)
	//将channel传入工人池,第一个64是job的数量
	worker_pool(64, send_channel, result_channel)
	//接收结果并打印
	go func(result_c chan *Result) {
		for result := range result_c {
			fmt.Println("第", result.job.Id+1, "个job,它的随机值是:", result.job.Random_number, "是它的结果是:", result.sum)
		}
	}(result_channel)
	//defer关闭channel
	defer close(send_channel)
	defer close(result_channel)
	//建立将job传入channel的函数(正常来说这一步应该是放在前面,但是文档这里是无限循环,所以放在最后,我这里改成有限循环了,注意这里循环值也就是循环次数,i的最大值要大一些,否则程序可能没来得及打印就推出了)
	for i := 0; i < 1000; i++ {
		rand_num := rand.Int()
		job := &Job{
			Id:            i,
			Random_number: rand_num,
		}
		send_channel <- job
	}
	// time.Sleep(2 * time.Second)
}

// 创建工人池
func worker_pool(num int, sc chan *Job, rc chan *Result) {
	for i := 0; i < num; i++ {
		go func(sc chan *Job, rc chan *Result) {
			for job := range sc {
				//获取随机数并进行处理
				r_num := job.Random_number
				//sum就是求和的结果,要把sum传进Result结构体,并将Result传入rc通道
				sum := 0
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num = r_num / 10
				}
				re := &Result{
					job: job,
					sum: sum,
				}
				rc <- re
			}
		}(sc, rc)
	}
}
