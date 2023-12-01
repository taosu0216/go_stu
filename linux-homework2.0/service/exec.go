package service

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

var Lines []string

// Login test handler
//
// @Tags	user
// @Summary		登陆
// @Description	Description 登陆
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json{"message"}
// @Router			/login [post]
func Login(ctx context.Context, c *app.RequestContext) {
	username := c.PostForm("username")
	passwd := c.PostForm("password")
	if username == "taosu" && passwd == "1433223" {
		c.JSON(200, utils.H{
			"msg": "登录成功",
		})
	}
}

// GetInfos 获取信息
//
//		@Tags	功能
//		@Description	获取信息
//		@Accept			application/json
//		@Produce		application/json
//	 	@Param Authorization header string false "Bearer JWT-TOKEN"
//		@Success 200 {string} json{"msg"}
//		@Router			/user/getinfos [get]
func GetInfos(ctx context.Context, c *app.RequestContext) {
	Sh()
	c.JSON(200, utils.H{
		"msg": Lines,
	})
}
func Sh() {
	cmd := exec.Command("sh", "-c", "./start.sh > infos.txt")
	if err := cmd.Run(); err != nil {
		log.Println("执行失败")
	}
	fmt.Println("执行成功")
	file, err := os.Open("./infos.txt")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		Lines = append(Lines, line)
	}
}
