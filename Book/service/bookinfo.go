package service

import (
	"Book/models"
	"Book/utils"
	"github.com/gin-gonic/gin"
	"strconv"
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

// UpdateBook
// @Summary	修改图书借阅状态
// @Tags    图书模块
// @Param 	id 	query 	string 	false 	"图书ID"
// @Param 	isreturn    query 	string 	false 	"图书状态"
// @Success	200	{string}	json{"code":"message"}
// @Router		/book/updatebook [get]
func UpdateBook(c *gin.Context) {
	book := &models.BookInfo{}
	idStr := c.Query("id")
	id, _ := strconv.ParseUint(idStr, 10, 0)
	book.ID = uint(id)
	book.IsReturn = c.Query("isreturn")
	IsFInd, books, _ := models.FindBookByID(book.ID)
	result := utils.DB.Model(&books).Updates(models.BookInfo{
		IsReturn: book.IsReturn,
	})
	if !IsFInd || result.Error != nil {
		c.JSON(400, gin.H{
			"message": "没找到该图书!",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "修改成功!",
		})
	}

}
