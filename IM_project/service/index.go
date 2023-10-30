package service

import (
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	//H 可以看作是 "Hash" 的缩写，通常用于表示具有键值对的数据结构
	c.JSON(200, gin.H{
		"message": "welecome!!",
	})
}
