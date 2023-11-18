package router

import (
	hz "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "memorandum/docs"
	"memorandum/service"
)

var identityKey = "username"

func RouterExec() *hz.Hertz {

	h := hz.Default(hz.WithHostPorts(":6789"))

	h.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	//用户注册,登陆模块
	h.POST("/signin", service.SignIn)
	h.POST("/login", RootMiddleware.LoginHandler, service.Login)

	url := swagger.URL("http://localhost:6789/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
	auth := h.Group("/auth")
	auth.Use(RootMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", service.Hello)
	}
	return h
}
