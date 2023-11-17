package utils

import (
	"JWT/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(user model.User) (tokenString string, err error) {
	expire := time.Now().Add(7 * 24 * time.Hour)
	claim := &model.MyClaims{
		UserId:   user.Id,
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),     //过期时间
			IssuedAt:  time.Now().Unix(), //发布时间
			Subject:   "token",           //主题
			Issuer:    "taosu",           //发布人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString(model.JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
