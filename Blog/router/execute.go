package router

import (
	"Blog/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	return r
}
