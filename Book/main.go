package main

import (
	"Book/router"
	"Book/utils"
	"log"
)

// @title 图书借阅系统
// @version 1.0
// @description Go练手项目
func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	err := r.Run("0.0.0.0:5678")
	if err != nil {
		log.Fatalln("running router err is ", err)
	}
}
