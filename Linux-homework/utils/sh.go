package utils

import (
	"fmt"
	"os/exec"
)

func Only() {
	mkdir := exec.Command("sh", "-c", "cd missions && mkdir AllInfos && cd AllInfos")
	if err := mkdir.Run(); err != nil {
		fmt.Println("创建文件夹失败")
	}
}
func Cpu_info_Sh() {
	cmd_info := exec.Command("sh", "-c", "cpu-info > ./missions/AllInfos/cpu_infos.txt")

	if err := cmd_info.Run(); err != nil {
		fmt.Println("获取cpu信息失败")
	}

	ReadCPU_Infos()
	fmt.Println("获取cpu信息成功")
}
func Mem_info_Sh() {
	memory_info := exec.Command("sh", "-c", "free -h > ./missions/AllInfos/memory_infos.txt")
	if err := memory_info.Run(); err != nil {
		fmt.Println("获取内存信息失败")
	}
	ReadMenory_Infos()
	fmt.Println("获取内存信息成功")
}
