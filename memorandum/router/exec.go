package router

import (
	hz "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
	"memorandum/MiddleWare"
	_ "memorandum/docs"
	"memorandum/service"
)

func EnterExec() *hz.Hertz {
	//jwt初始化
	JwtInit()

	//框架,swagger初始化
	h := hz.Default(hz.WithHostPorts(":6789"))
	h.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	url := swagger.URL("http://localhost:6789/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	//用户模块
	{
		//注册
		h.POST("/signin", service.SignIn)
		//登陆
		h.POST("/login", MiddleWare.UserMiddleware.LoginHandler, service.Login)
	}

	//测试模块
	auth := h.Group("/auth")
	auth.Use(MiddleWare.RootMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", service.Hello)
	}

	//用户功能模块
	user := h.Group("/user")
	user.Use(MiddleWare.UserMiddleware.MiddlewareFunc())
	{
		user.GET("/test", service.Test)
		user.GET("/createitem", service.CreateItem)
		user.GET("/finditem", service.FindItem)
		user.POST("/edititemstatus", service.EditItemStatus)
		user.POST("/deleteitem", service.DeleteItem)
	}
	return h
}
func JwtInit() {
	//测试管理员
	MiddleWare.RootJwtInit()
	//用户鉴权
	MiddleWare.UserJwtInit()
}
