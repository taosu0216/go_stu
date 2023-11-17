package models

import (
	"Book/handler"
	"Book/utils"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"time"
)

func CreateUser(user *UserInfo) bool {
	result := utils.DB.Create(user)
	if result.Error != nil {
		isErr := handler.Handler(result.Error, "create user failed...")
		return isErr
	}
	return false
}

func FindUserByName(name string) (bool, *UserInfo, *gorm.DB) {
	user := &UserInfo{}
	data := utils.DB.Where("name = ?", name).First(user)
	if user.Name != "" {
		return true, user, data
	}
	return false, user, data
}

func IsUserExit(user *UserInfo) bool {
	isExit, _, _ := FindUserByName(user.Name)
	if isExit {
		return true
	} else {
		return false
	}
}

// MakeToken 建立token
func MakeToken(u *UserInfo) (string, error) {
	expirationTime := time.Now().Add(3 * time.Minute)
	claims := &MyClaims{
		UserId:   u.Id,
		Username: u.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   "token",
			Issuer:    "taosu",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	return token, claims, err
}
