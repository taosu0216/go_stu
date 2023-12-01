package utils

import (
	"fmt"
)

func ErrIsExist(err error, msg string) bool {
	if err != nil {
		fmt.Println(msg, err)
		return true
	}
	return false
}
