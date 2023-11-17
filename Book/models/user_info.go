package models

import (
	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("1433223qluorangestudiocn_SBCMBDSJNabjBKBKBUHKLJbkjds")

type UserInfo struct {
	//gorm.Model
	Id       int    `gorm:"primarykey;autoIncrement" json:"id"`
	Name     string `json:"username"`
	PassWord string
	QQ       string
	Email    string `valid:"email"`
	Salt     string
}

type MyClaims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func (table *UserInfo) TableName() string {
	return "user_infos"
}
