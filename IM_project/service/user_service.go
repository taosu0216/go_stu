package service

import (
	"IM_project/models"
	"IM_project/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GetUserList
// @Summary	所有用户
// @Tags		用户模块
// @Success	200	{string}	json{"code":"message"}
// @Router		/user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	//输入账密,并重复输入密码,判断两次密码是否一致
	user := models.UserBasic{}
	//这里的c.Query是表单参数,也就是url参数
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	//判断是否名字是否被占用
	err_isExit, isExit := models.IsExit(user)
	if isExit {
		c.JSON(400, gin.H{
			"message": err_isExit,
		})
		return
	}
	//400是bad request,即请求失败
	if password != repassword {
		c.JSON(400, gin.H{
			"message": "两次密码不一致!",
		})
		return
	}
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	err := models.CreateUser(user)
	if err != nil {
		log.Fatalln(err)
	}
	_ = models.UpdateLoginTime(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "用户创建成功",
		"data": gin.H{
			"name":       user.Name,
			"password":   password,
			"login_time": user.LoginTime,
		},
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @Param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	//c.Query()查询的结果返回值是string,用strconv()来转换成int型
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Fatalln(err)
	}
	user.ID = uint(id)

	//判断要删除的用户是否存在
	realid, err := models.DeleteUser(user)
	if err != nil {
		log.Fatalln(err)
	}
	if realid == 0 {
		c.JSON(400, gin.H{
			"message": "用户不存在",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "用户删除成功",
	})
}

// UpdateUser
// @Summary 修改用户信息
// @Tags 用户模块
// @Param id formData string false "id"
// @Param name formData string false "用户名"
// @Param password formData string false "密码"
// @Param email formData string false "邮箱"
// @Param phone formData string false "手机号"
// @Success 200 {string} json{"code","message","data"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	//根据用户id修改信息,但是id不存在时未进行校验
	user := models.UserBasic{}
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		log.Fatalln(err)
	}
	user.ID = uint(id)
	is_Exit, _ := models.FindUserById(user.ID)
	if !is_Exit {
		c.JSON(400, gin.H{
			"message": "用户不存在!",
		})
		return
	}
	user.Name = c.PostForm("name")
	passwd := c.PostForm("password")
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.PassWord = utils.MakePassword(passwd, salt)
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	user.Salt = salt
	//对用户输入信息进行校验
	//调用govalidator包的ValidateStruct方法对传入的结构体进行内容校验,校验方法就是在定义结构体时`valid:""`里写的内容
	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"message": "修改用户信息格式错误!",
		})
		return
	} else {
		_ = models.UpdateLoginTime(user)
		err = models.UpdateUser(user)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("update :", user)
		c.JSON(200, gin.H{
			"code":    0,
			"message": "修改用户信息成功",
			"data": gin.H{
				"name":       user.Name,
				"password":   passwd,
				"phone":      user.Phone,
				"email":      user.Email,
				"login_time": user.LoginTime,
			},
		})
	}
}

// Login
// @Summary 用户登录
// @Tags 用户模块
// @Param name formData string false "用户名"
// @Param password formData string false "密码"
// @Success 200 {string} json{"code","message","data"}
// @Router /user/login [post]
func Login(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	isExit, userDB, user_data := models.FindUserByName(user.Name)
	if !isExit {
		c.JSON(200, gin.H{
			"message": "该用户不存在!",
		})
		return
	}
	if !utils.ValidPassword(user.PassWord, user_data.Salt, user_data.PassWord) {
		c.JSON(200, gin.H{
			"message": "密码错误!",
		})
		return
	}
	refresh_token := utils.MakeToken()
	models.FindUserAndRefreshToken(userDB, refresh_token)
	_ = models.UpdateLoginTime(user_data)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "登录成功!",
		"data": gin.H{
			"name":       user_data.Name,
			"phone":      user_data.Phone,
			"email":      user_data.Email,
			"login_time": user_data.LoginTime,
		},
	})
}

//redis

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("close:", err)
		}
	}(conn)
	MsgHandler(conn, c)
}
func MsgHandler(conn *websocket.Conn, c *gin.Context) {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if err != nil {
		fmt.Println(err)
	}
	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[conn][%s] : %s", tm, msg)
	err = conn.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println(err)
	}
}
