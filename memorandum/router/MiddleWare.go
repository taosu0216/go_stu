package router

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
	RootMiddleware *jwt.HertzJWTMiddleware
)

func Root() {
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
			claims := jwt.ExtractClaims(ctx, c)
			return &model.UserInfo{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			name := c.PostForm("username")
			passwd := c.PostForm("password")
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
			if v, ok := data.(*model.UserInfo); ok && v.Username == "taosu" {
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
