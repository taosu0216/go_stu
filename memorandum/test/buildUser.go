package test

import (
	"fmt"
	"memorandum/model"
	"memorandum/util"
)

func BuildUserTable() {
	err := util.DB.AutoMigrate(&model.UserInfo{})
	util.Err(err, "数据库迁移失败...")
	fmt.Println("数据库迁移成功...")
}
