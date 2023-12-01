package utils

import (
	"Linux-homework/model"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var mem model.Memory

func ReadMenory_Infos() {
	file, err := os.Open("./AllInfos/memory_infos.txt")
	if !ErrIsExist(err, "打开系统内存信息文件失败") {
		fmt.Println("打开系统内存信息文件成功")
		defer file.Close()
		scanner := bufio.NewScanner(file)
		Total, Used, Free, Shared, Buff, Available := GetAll(scanner)
		mem.Total = Total
		mem.Used = Used
		mem.Free = Free
		mem.Shared = Shared
		mem.Buff = Buff
		mem.Available = Available
		if create_err := DB.Create(&mem).Error; create_err != nil {
			log.Fatalln("内存信息插入失败")
		} else {
			fmt.Println("内存信息插入成功")
		}
	}
}
func GetAll(scanner *bufio.Scanner) (string, string, string, string, string, string) {
	var Total, Used, Free, Shared, Buff, Available string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Mem:") {
			fileds := strings.Fields(line)
			if len(fileds) >= 7 {
				Total = fileds[1]
				Used = fileds[2]
				Free = fileds[3]
				Shared = fileds[4]
				Buff = fileds[5]
				Available = fileds[6]
			}
		}
	}
	return Total, Used, Free, Shared, Buff, Available
}
