package main

import (
	"Linux-homework/utils"
)

func main() {
	utils.InitMysql()
	utils.Install()
	utils.BuildTable()
	utils.Only()
	utils.Sh()
}
