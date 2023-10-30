package router

import (
	"IM_project/docs"
	"IM_project/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	//创建基本路由
	r := gin.Default()
	//这里必须访问swagger/index.html而不是swagger/index
	/*
		这里的SwaggerInfo.BasePath = "" 是指根路径为空
		即访问时直接127.0.0.1:8099/swagger/index.html即可,
		如果不为空,假设是SwaggerInfo.BasePath = "test"的话
		那就是访问127.0.0.1:8099/test/swagger/index.html
	*/
	docs.SwaggerInfo.BasePath = ""
	/*
		这里的*any是自定义的通配符的名字,就是swagger/下所有请求都交由这里的ginSwagger.WrapHandler(swaggerfiles.Handler)来处理
		这里的ginSwagger.WrapHandler(swaggerfiles.Handler)也是固定用法,记着就行
	*/
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//处理方式放在service中处理,这里不是service.GetIndex()
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/user/createUser", service.CreateUser)
	r.GET("/user/deleteUser", service.DeleteUser)
	r.POST("/user/login", service.Login)
	r.POST("/user/updateUser", service.UpdateUser)

	r.GET("/user/SendMsg", service.SendMsg)
	return r
}
