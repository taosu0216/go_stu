package utils

import (
	"fmt"
	"os/exec"
)

func Only() {
	mkdir := exec.Command("sh", "-c", "mkdir AllInfos&&cd AllInfos")
	if err := mkdir.Run(); err != nil {
		fmt.Println("创建文件夹失败")
	}
}
func Sh() {

	cmd_info := exec.Command("sh", "-c", "cpu-info >> ./AllInfos/cpu_infos.txt")
	memory_info := exec.Command("sh", "-c", "free -h >> ./AllInfos/memory_infos.txt")

	if err := cmd_info.Run(); err != nil {
		fmt.Println("获取cpu信息失败")
	}
	if err := memory_info.Run(); err != nil {
		fmt.Println("获取内存信息失败")
	}
	ReadCPU_Infos()
	ReadMenory_Infos()
	fmt.Println("获取信息成功")
}
