package main

import (
	"memorandum/router"
	"memorandum/test"
	"memorandum/util"
)

//	@title			HertzTest
//	@version		1.0
//	@description	This is a demo using Hertz.

func main() {
	util.InitConfig()
	util.InitMysql()
	test.BuildUserTable()
	h := router.RouterExec()
	h.Spin()
}
