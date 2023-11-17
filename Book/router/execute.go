package router

import (
	"Book/docs"
	"Book/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	//注册路由
	r := gin.Default()
	r.Use(service.Origin())
	//加载静态资源
	r.LoadHTMLGlob("static/html/*")
	r.Static("/static", "./static")

	//配置swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//图书展示界面
	r.GET("/index", service.Index)

	//root用户能干的事
	//图书的CRUD
	//查看所有图书
	r.GET("/book/booklists", service.MiddleWare(), service.GetBookLists)
	//添加图书
	r.GET("/book/addbook", service.MiddleWare(), service.AddBook)
	//修改图书借阅状态(根据图书ID查询)
	r.GET("/book/updatebook", service.MiddleWare(), service.UpdateBook)

	//展示登陆页面(Get请求),当登陆时使用Post请求,方便管理只允许root用户修改图书信息
	r.GET("/user/login", service.MiddleWare(), service.ShowLoginPage)
	r.POST("/user/login", service.MiddleWare(), service.Login)
	//当有新的人来管理图书时,填写新的用户信息
	r.POST("/user/signin", service.MiddleWare(), service.CreateUser)

	//管理员修改图书信息前端界面
	r.GET("/user/index", service.MiddleWare(), service.ShowIndexPage)
	return r
}
