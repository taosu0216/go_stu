package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	UserImg  string `json:"user_img"`
}
type Todo struct {
	gorm.Model
	UserId uint   `json:"user_id"`
	Item   string `json:"item"`
	Status string `json:"status"`
}

func (table *UserInfo) TableName() string {
	return "user_infos"
}
func (table *Todo) TableName() string {
	return "todos"
}
