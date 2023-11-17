package middleware

import (
	"JWT/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *model.MyClaims, error) {
	claims := &model.MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return model.JwtKey, nil
	})
	return token, claims, err
}

// AuthMiddleware 定义中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		tokeString := c.GetHeader("token")
		fmt.Println(tokeString, "当前token")
		if tokeString == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "必须传递token",
			})
			c.Abort()
			return
		}
		token, claims, err := ParseToken(tokeString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "token解析错误",
			})
			c.Abort()
			return
		}
		// 从token中解析出来的数据挂载到上下文上,方便后面的控制器使用
		c.Set("userId", claims.UserId)
		c.Set("userName", claims.Username)
		c.Next()
	}
}
