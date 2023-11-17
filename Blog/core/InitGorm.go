package core

import (
	"Blog/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConn() *gorm.DB {
	dsn := global.Config.Mysql.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.Log.Error("failed to connect MySql")
		return nil
	}
	// 迁移 schema
	//err = db.AutoMigrate(&Product{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	fmt.Println("Mysql Connected successfully...")
	return db
}
