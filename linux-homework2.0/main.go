package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	hz "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/jwt"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "linux-homework2.0/docs"
	"linux-homework2.0/service"
)

var (
	UserMiddleware  *jwt.HertzJWTMiddleware
	identityKey     = "username"
	useridentityKey = "username"
)

type User struct {
	UserName string
	Password string
}

func main() {
	h := hz.Default(hz.WithHostPorts(":12345"))
	h.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	url := swagger.URL("http://localhost:12345/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
	UserJwtInit()
	h.POST("/login", UserMiddleware.LoginHandler, service.Login)
	user := h.Group("/user")
	user.Use(UserMiddleware.MiddlewareFunc())
	{
		user.GET("/getinfos", service.GetInfos)
	}
	h.Spin()
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
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
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
			if name == "taosu" && passwd == "1433223" {
				return &User{
					UserName: name,
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
	errInit := UserMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatalln("AuthMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}
