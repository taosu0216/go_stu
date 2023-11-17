package service

import (
	"Book/models"
	"Book/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

// CreateUser
// @Summary	注册用户
// @Tags    用户模块
// @Param	name		formData	string	false	"用户名"
// @Param 	password 	formData 	string 	false 	"密码"
// @Param 	repassword 	formData 	string 	false 	"重新确认密码"
// @Param 	qq 			formData 	string 	false 	"QQ号"
// @Param   email 		formData 	string 	false 	"邮箱"
// @Success	200	{string}	json{"code":"message"}
// @Router		/user/signin [post]
func CreateUser(c *gin.Context) {
	user := &models.UserInfo{}
	user.Name = c.PostForm("name")
	passwd := c.PostForm("password")
	repasswd := c.PostForm("repassword")
	user.QQ = c.PostForm("qq")
	user.Email = c.PostForm("email")
	user.Salt = fmt.Sprintf("%06d", rand.Int31())
	//判断用户名是否被占用
	if models.IsUserExit(user) {
		c.JSON(400, gin.H{
			"message": "用户名已存在!",
		})
		return
	}
	//判断密码是否一致
	if passwd != repasswd {
		c.JSON(400, gin.H{
			"message": "两次密码不一致!",
		})
		return
	}
	user.PassWord = utils.MakePasswd(passwd, user.Salt)
	token, err := models.MakeToken(user)
	if err != nil || token == "" {
		c.JSON(400, gin.H{
			"message": "token生成失败!",
		})
		return
	}
	isFind := models.CreateUser(user)
	if isFind {
		c.JSON(400, gin.H{
			"message": "failed to create user...",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "用户创建成功!",
			"token":   token,
		})
	}
}

// Login
// @Summary 用户登录
// @Tags 用户模块
// @Param name formData string false "用户名"
// @Param password formData string false "密码"
// @Success 302 {string} json{"message","data"}
// @Router /user/login [post]
func Login(c *gin.Context) {
	user := &models.UserInfo{}
	user.Name = c.PostForm("name")
	password := c.PostForm("password")
	NameIsFind, user, _ := models.FindUserByName(user.Name)
	if !NameIsFind {
		c.JSON(400, gin.H{
			"message": "用户名不存在!",
		})
	} else if !utils.ValidPasswd(password, user.Salt, user.PassWord) {
		c.JSON(400, gin.H{
			"message": "密码错误!",
		})
	} else {
		token, err := models.MakeToken(user)
		if err != nil || token == "" {
			c.JSON(400, gin.H{
				"message": "token生成失败!",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "登录成功,跳转首页",
			"name":    user.Name,
			"qq":      user.QQ,
			"token:":  token,
		})
	}
}
