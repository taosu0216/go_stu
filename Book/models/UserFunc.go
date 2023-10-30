package models

import (
	"Book/handler"
	"Book/utils"
	"gorm.io/gorm"
)

func GetUserLists() []*UserInfo {
	lists := make([]*UserInfo, 300)
	utils.DB.Find(&lists)
	return lists
}
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

func UserInfoUpdate(user *UserInfo) error {
	result := utils.DB.Model(&user).Updates(UserInfo{
		Token: user.Token,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
