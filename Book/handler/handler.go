package handler

import "fmt"

func Handler(err error, msg string) bool {
	if err != nil {
		fmt.Println("error message is : ", msg, "error is : ", err)
		return true
	}
	return false
}
