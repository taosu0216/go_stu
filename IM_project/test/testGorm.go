package main

import (
	"IM_project/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Code  string
	Price uint
}

func main() {
	//"用户名:密码@tcp(地址:端口)/数据库名"
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/IM?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	//这里连接IM数据库后,会自动使用结构体名作为表名(按理说是全变小写+复数,但是我看的没加复数,只变成小写了)
	user := &models.UserBasic{}
	user.Name = "taosu"
	db.Create(user)

	// Read
	fmt.Println("db.First(user, 1) : ", db.First(user, 1))

	// Update - 将 user 的 password 更新为 1234
	db.Model(user).Update("PassWord", "1234")
}
