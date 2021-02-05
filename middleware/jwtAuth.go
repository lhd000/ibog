package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"iblog/service"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.GetHeader("Authorization")

		//validate toke formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusOK, gin.H{"code": 401, "msg": "please login first"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		userinfo,err := service.ParseToken(tokenString)

		if err != nil {

			c.JSON(http.StatusOK, gin.H{"code": 401, "msg": "validate fail"})
			c.Abort()
			return
		}


		//用户存在 将user的信息写入上下文
		c.Set("user",userinfo)
		c.Next()
	}
}