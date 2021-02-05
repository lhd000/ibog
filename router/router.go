package router

import (
	"iblog/config"
	"iblog/controller"
	"iblog/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func RegistRouter(this *gin.Engine) {

	store := cookie.NewStore([]byte(config.COOKIE_SECRET_STR))
	this.Use(sessions.Sessions(config.SESSION_STORE_NAME, store))

	admin := this.Group("/admin")
	admin.Use(middleware.CkLogin())
	{
		admin.GET("/users", controller.GetUserList)

		admin.GET("/index", controller.Admin)

		admin.GET("/", func(ctx *gin.Context) {
			ctx.Redirect(302, "/admin/index")
		})

		admin.GET("/article", controller.Article)

		admin.GET("/category", controller.AdminCategory)

		admin.POST("/categoryadd", controller.CategoryAdd)

		admin.GET("/categoryupdate", controller.Categoryupdate)

		admin.GET("/categorydel", controller.CategoryDel)

		admin.GET("/articleadd", controller.ArticleAdd)

		admin.POST("/articleadd", controller.ArticleAddDo)

		admin.GET("/articledelete", controller.ArticleDelete)

		admin.GET("/articleupdate", controller.ArticleUpdatePage)

		admin.POST("/articleupdate", controller.ArticleUpdate)

		admin.GET("/manageuser", controller.Manageuser)

		admin.POST("/useradd", controller.UserAdd)

		admin.GET("/deluser", controller.UserDel)
		admin.GET("/logout", controller.LogOut)

	}

	this.Any("/ueditor", controller.Action)

	this.GET("/front", controller.Front)

	this.Static("/static", "./static")

	this.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/front")

	})

	this.GET("/front/aboutus", controller.AboutUs)
	this.GET("/front/contacts", controller.Contacts)
	this.GET("/front/category", controller.FrontCategory)
	this.GET("/front/author", controller.Author)
	this.GET("/front/blogpost", controller.BlogPost)
	this.GET("/front/blank", controller.Blank)

	this.GET("/login", controller.Login)
	this.POST("/login", controller.LoginAction)
}
