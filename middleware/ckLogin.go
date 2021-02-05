package middleware

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//Cros  fix the cros problem

func CkLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		u := session.Get("admin")

		log.Println("the login info is :", u)

		if u == nil {
			c.Redirect(302, "/login")
			return
		}

		c.Next()
	}
}
