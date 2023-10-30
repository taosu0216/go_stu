package models

import (
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Name     string `gorm:"foreignKey:UserName"`
	PassWord string
	QQ       string
	Email    string `valid:"email"`
	Phone    string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Token    string
	Salt     string
}

type Borrowed struct {
	gorm.Model
	UserName         string
	BorrowedBookName string
	BorrowedTime     string
	ReturnTime       string
	IsReturn         string
}

func (table *UserInfo) TableName() string {
	return "user_infos"
}
