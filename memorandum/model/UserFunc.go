package model

import (
	"memorandum/util"
)

func FindUserByName(name string) (bool, *UserInfo) {
	user := &UserInfo{}
	_ = util.DB.Where("username = ?", name).First(user)
	if user.Username != "" {
		return true, user
	}
	return false, nil
}
func FindIdByName(name interface{}) (bool, uint) {
	user := &UserInfo{}
	_ = util.DB.Where("username = ?", name).First(user)
	if user.Username != "" {
		return true, user.ID
	}
	return false, 0
}
func FindItemsByUsername(username interface{}) (bool, []Todo) {
	isExist, userId := FindIdByName(username)
	if isExist {
		//var items []string
		var todos []Todo
		//_ = util.DB.Table("todos").Where("user_id = ?", userId).Pluck("item", &items)
		_ = util.DB.Table("todos").Where("user_id = ?", userId).Find(&todos)
		return true, todos
	}
	return false, nil
}
func FindTodoUserByItemAndUserId(item string, user_id uint) (bool, *Todo) {
	todo := &Todo{}
	_ = util.DB.Where("item = ? AND user_id = ?", item, user_id).First(todo)
	if todo.UserId != 0 && todo.Item != "" {
		return true, todo
	}
	return false, nil
}
