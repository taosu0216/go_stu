package service

import (
	"Book/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	lists := models.GetBookLists()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "C109图书借阅系统",
		"books": lists,
	})
}
func Userinfo(c *gin.Context) {
	c.HTML(http.StatusOK, "info.html", gin.H{
		"title": "用户主界面",
	})
}
func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "C109图书角",
	})
}
