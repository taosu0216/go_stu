package router

import (
	mid "JWT/middleware"
	"JWT/model"
	"JWT/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var loginStruct model.LoginStruct
		err := c.BindJSON(&loginStruct)
		if err != nil {
			fmt.Println("解析数据错误")
			c.JSON(200, gin.H{
				"code":    1,
				"message": "数据解析错误",
			})
			return
		}
		fmt.Println("登陆参数", loginStruct)
		userInfo := model.User{
			Id:       "1",
			UserName: loginStruct.Username,
		}
		token, err := utils.GenerateToken(userInfo)
		if err != nil {
			fmt.Println(err, "============")
			c.JSON(200, gin.H{
				"code":    1,
				"message": "生成token错误",
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    0,
			"message": "登陆成功",
			"token":   token,
		})

	})
	router.POST("/check_token", mid.AuthMiddleware(), func(c *gin.Context) {
		if userName, exists := c.Get("userName"); exists {
			fmt.Println("当前用户", userName)
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "验证成功",
		})
	})
	return router
}
