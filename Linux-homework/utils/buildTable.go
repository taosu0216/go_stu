package utils

import (
	"Linux-homework/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	dsn := "root:root@tcp(127.0.0.1:3306)/Linuxhomework?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("mysql连接失败......")
	}
	fmt.Println("mysql连接成功......")
}
func BuildCPUtable() {
	err := DB.AutoMigrate(&model.CPU{})
	if !ErrIsExist(err, "cpu表创建失败") {
		fmt.Println("cpu信息表创建成功")
	}
}
func BuildMemoryTable() {
	err := DB.AutoMigrate(&model.Memory{})
	if !ErrIsExist(err, "内存信息表创建失败") {
		fmt.Println("内存信息表创建成功")
	}
}
func BuildTable() {
	BuildCPUtable()
	BuildMemoryTable()
}
