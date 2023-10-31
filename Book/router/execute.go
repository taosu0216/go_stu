package router

import (
	"Book/docs"
	"Book/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("static/html/*")
	r.Static("/static", "./static")
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/book/booklists", service.GetBookLists)
	r.GET("/userlists", service.GetUserLists)
	r.POST("/user/signin", service.CreateUser)
	r.GET("/user/login", service.ShowLoginPage)
	r.POST("/user/login", service.Login)
	r.GET("/book/addbook", service.AddBook)
	r.GET("/user/index.html", service.Index)
	r.GET("/user/info.html", service.Userinfo)
	return r
}
