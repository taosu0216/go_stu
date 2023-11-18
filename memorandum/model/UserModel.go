package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	UserImg  string `json:"user_img"`
}

func (table *UserInfo) TableName() string {
	return "user_infos"
}
