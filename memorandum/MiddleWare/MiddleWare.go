package MiddleWare

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"log"
	"memorandum/model"
	"time"
)

var (
	RootMiddleware  *jwt.HertzJWTMiddleware
	UserMiddleware  *jwt.HertzJWTMiddleware
	identityKey     = "username"
	useridentityKey = "username"
)

func RootJwtInit() {
	var err error
	RootMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "冻干柠檬汁",
		Key:         []byte("14332232022850100"),
		Timeout:     time.Hour * 72,
		MaxRefresh:  time.Hour * 48,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Println(1234)
			if v, ok := data.(*model.UserInfo); ok {
				fmt.Println(v)
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			fmt.Println("到了吗")
			claims := jwt.ExtractClaims(ctx, c)
			fmt.Println(111)
			return &model.UserInfo{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			name := c.PostForm("username")
			passwd := c.PostForm("password")
			fmt.Println(name, passwd, "这里")
			isExit, user := model.FindUserByName(name)
			if isExit && user.Password == passwd {
				return &model.UserInfo{
					Username: name,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},

		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			fmt.Println("is ok?")
			if v, ok := data.(*model.UserInfo); ok && v.Username == "taosu" || v.Username == "admin" {
				fmt.Println("鉴权成功")
				return true
			}

			return false
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			fmt.Println("认证失败")
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},
		//TokenLookup:   "header: Authorization",
		TokenHeadName: "taosu",
		//TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatalln("JWT Error:" + err.Error())
	}
	errInit := RootMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatalln("AuthMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}
func UserJwtInit() {
	var err error
	UserMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "冻干柠檬汁",
		Key:         []byte("14332232022850100"),
		Timeout:     time.Hour * 72,
		MaxRefresh:  time.Hour * 48,
		IdentityKey: useridentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.UserInfo); ok {
				return jwt.MapClaims{
					useridentityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[useridentityKey].(string)
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			name := c.PostForm("username")
			passwd := c.PostForm("password")
			//fmt.Println("name: ", name, "passwd: ", passwd)
			isExit, user := model.FindUserByName(name)
			if isExit && user.Password == passwd {
				return &model.UserInfo{
					Username: name,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},

		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			fmt.Println("认证失败")
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},
		//TokenLookup:   "header: Authorization",
		TokenHeadName: "taosu",
		//TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatalln("JWT Error:" + err.Error())
	}
	errInit := RootMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatalln("AuthMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}
