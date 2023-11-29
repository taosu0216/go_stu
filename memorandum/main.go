package main

import (
	"memorandum/router"
	"memorandum/test"
	"memorandum/util"
)

//	@title			Taosuの备忘录demo
//	@version		1.0
//	@description	用Hertz创建的简单备忘录demo.

func main() {
	util.InitConfig()
	util.InitMysql()
	test.BuildUserTable()
	h := router.EnterExec()
	h.Spin()
}
