package buildUserTable

import (
	"Book/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func BuildUser() {
	//"用户名:密码@tcp(地址:端口)/数据库名"
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	_ = db.AutoMigrate(&models.UserInfo{})
	// Create
	//这里连接book数据库后,会自动使用结构体名作为表名(按理说是全变小写+复数,但是我看的没加复数,只变成小写了)
	user := &models.UserInfo{}
	user.Name = "test"
	db.Create(user)
	//db.Model(user).Update("PassWord", "root")
	fmt.Println("user info create success...")
}
