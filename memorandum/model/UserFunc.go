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
