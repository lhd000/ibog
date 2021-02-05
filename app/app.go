package app

import (
	"iblog/config"
	"iblog/middleware"
	"iblog/router"

	"github.com/gin-gonic/gin"
)

/**
  应用启动函数
*/
func Run() {

	r := gin.Default()

	//设置模板目录
	r.LoadHTMLGlob(config.GIN_TEMPLATE_DIR)
	//r.LoadHTMLFiles(config.GIN_TEMPLATE_DIR)

	//设置引用资源位置
	r.Static("/statics", "static")

	r.Use(middleware.Cors())

	router.RegistRouter(r)

	//设置应用启动 端口
	r.Run(":" + config.GIN_PORT) // listen and serve on 0.0.0.0:8080

}
