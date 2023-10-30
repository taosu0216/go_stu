package service

import (
	"Book/models"
	"Book/utils"
	"github.com/gin-gonic/gin"
)

// GetBookLists
// @Summary	图书列表
// @Tags    图书模块
// @Success	200	{string}	json{"code":"message"}
// @Router		/book/booklists [get]
func GetBookLists(c *gin.Context) {
	lists := models.GetBookLists()
	c.JSON(200, gin.H{
		"booklists": lists,
	})
}

// AddBook
// @Summary	添加图书
// @Tags    图书模块
// @Param 	bookname 	query 	string 	false 	"书名"
// @Param 	author	query 	string 	false 	"作者"
// @Param 	img 	query	string 	false 	"封面图片"
// @Param 	desc 	query   string 	false 	"描述"
// @Param 	tag 	query   string 	false 	"tag"
// @Param 	isreturn    query 	string 	false 	"图书状态"
// @Success	200	{string}	json{"code":"message"}
// @Router		/book/addbook [get]
func AddBook(c *gin.Context) {
	book := &models.BookInfo{}
	book.Name = c.Query("bookname")
	book.Author = c.Query("author")
	book.Img = c.Query("img")
	book.Tags = c.Query("tag")
	book.Description = c.Query("desc")
	book.IsReturn = c.Query("isreturn")
	utils.DB.Create(book)
	c.JSON(200, gin.H{
		"message": "图书添加成功!",
	})
}
