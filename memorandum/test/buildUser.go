package test

import (
	"fmt"
	"memorandum/model"
	"memorandum/util"
)

func BuildUserTable() {
	err := util.DB.AutoMigrate(&model.UserInfo{})
	util.Err(err, "数据库-用户表 迁移失败...")
	err = util.DB.AutoMigrate(&model.Todo{})
	util.Err(err, "数据库-待办事项表 迁移失败...")
	fmt.Println("数据库迁移成功...")
}
