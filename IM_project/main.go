package main

import (
	"IM_project/router"
	"IM_project/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := router.Router()
	r.Run("127.0.0.1:8099")
}
