package utils

import (
	"Linux-homework/model"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var cpu model.CPU

func ReadCPU_Infos() {
	file, err := os.Open("./AllInfos/cpu_infos.txt")
	if !ErrIsExist(err, "打开CPU信息文件失败") {
		fmt.Println("打开CPU信息文件成功")
		defer file.Close()
		scanner := bufio.NewScanner(file)
		packages := GetPackages(scanner)
		microarchitectures := GetMicroarchitectures(scanner)
		corecount := GetCoreCount(scanner)
		cpu.Packages = packages
		cpu.CoreCount = corecount
		cpu.Microarchitectures = microarchitectures
		if create_err := DB.Create(&cpu).Error; create_err != nil {
			log.Fatalln("cpu信息插入失败")
		} else {
			fmt.Println("cpu信息插入成功")
		}
	}
}
func GetPackages(scanner *bufio.Scanner) string {
	flag := false
	isPackagesLine := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Packages:") {
			isPackagesLine = true
		}
		if isPackagesLine && flag {
			packages := line
			fmt.Println("packages is : ", packages)
			flag = false
			isPackagesLine = false
			return packages
		}
		if isPackagesLine {
			flag = true
		}
	}
	return ""
}
func GetMicroarchitectures(scanner *bufio.Scanner) string {
	flag := false
	isMicroarchitecturesLine := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Microarchitectures:") {
			isMicroarchitecturesLine = true
		}
		if isMicroarchitecturesLine && flag {
			microarchitectures := line
			fmt.Println("microarchitectures is : ", microarchitectures)
			flag = false
			isMicroarchitecturesLine = false
			return microarchitectures
		}
		if isMicroarchitecturesLine {
			flag = true
		}
	}
	return ""
}
func GetCoreCount(scanner *bufio.Scanner) int {
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "APIC") {
			count++
		}
	}
	return count
}
