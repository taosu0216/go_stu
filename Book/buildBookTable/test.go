package buildBookTable

import (
	"Book/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func BuildBook() {
	//"用户名:密码@tcp(地址:端口)/数据库名"
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	_ = db.AutoMigrate(&models.BookInfo{})
	// Create
	//这里连接book数据库后,会自动使用结构体名作为表名(按理说是全变小写+复数,但是我看的没加复数,只变成小写了)
	book := &models.BookInfo{}
	book.Name = "web编程实战"
	book.IsReturn = "未归还"
	book.BorrowedTime = time.Now().Format("2006-01-02 15:04:05")
	book.BorrowerName = "root"
	db.Create(book)
	//db.Model(book).Update("BookId", "1")
	fmt.Println("book info create success...")
}
