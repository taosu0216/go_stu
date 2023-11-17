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
func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "C109图书角",
	})
}

func ShowIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "userindex.html", gin.H{
		"title": "C109图书角",
	})
}
