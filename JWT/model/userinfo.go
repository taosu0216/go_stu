package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
}

type MyClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// LoginStruct 登录的参数
type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JwtKey 证书签名密钥
var JwtKey = []byte("1433223")
