package service

import (
	"Book/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.Next()
			return
		}
		token, _, err := models.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			c.AbortWithStatus(403)
			return
		}
		c.Redirect(302, "index")
		c.Abort()
	}
}

// Origin 跨域设置
func Origin() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		//defer func() {
		//	if err := recover(); err != nil {
		//		log.Logger.Error("HttpError", zap.Any("HttpError", err))
		//	}
		//}()

		c.Next()
	}
}
