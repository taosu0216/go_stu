package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"memorandum/model"
	"memorandum/util"
)

// Hello 123 handler
//
//		@Summary		你好
//		@Description	Description 你好
//		@Accept			application/json
//		@Produce		application/json
//	 @Param Authorization header string false "Bearer JWT-TOKEN"
//		@Success 200 {string} json{"data","msg"}
//		@Router			/auth/hello [get]
func Hello(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, utils.H{
		"data": "hello",
		"msg":  "world",
	})
}

// Login test handler
//
// @Summary		登陆
// @Description	Description 登陆
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json{"message"}
// @Router			/login [post]
func Login(ctx context.Context, c *app.RequestContext) {
	username := c.PostForm("username")
	passwd := c.PostForm("password")
	isExit, user := model.FindUserByName(username)
	if !isExit {
		c.JSON(200, utils.H{
			"message": "用户不存在",
		})
		return
	}
	if user.Password != passwd {
		c.JSON(200, utils.H{
			"message": "密码错误",
		})
		return
	}
	c.JSON(200, utils.H{
		"message": "登陆成功",
	})
}

// SignIn test handler
// @Summary		注册
// @Description	Description 注册
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json{"message"}
// @Router			/signin [post]
func SignIn(ctx context.Context, c *app.RequestContext) {
	user := &model.UserInfo{}
	name := c.PostForm("username")
	password := c.PostForm("password")
	isExit, _ := model.FindUserByName(name)
	if isExit {
		c.JSON(200, utils.H{
			"message": "用户名已被注册",
		})
		return
	}
	user.Username = name
	user.Password = password
	util.DB.Create(&user)
	c.JSON(200, utils.H{
		"message": "注册成功",
	})
}
