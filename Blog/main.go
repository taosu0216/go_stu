package main

import (
	"Blog/core"
	"Blog/global"
	"Blog/router"
	"fmt"
)

func main() {
	core.InitConfig()
	global.Log = core.InitLogger()
	global.DB = core.MysqlConn()
	r := router.InitRouter()
	addr := global.Config.System.AddrString()
	fmt.Println("Blog运行在: ", addr)
	err := r.Run(addr)
	if err != nil {
		global.Log.Error(err)
	}
}
