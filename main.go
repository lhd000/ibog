package main

import (
	// "github.com/gin-gonic/gin"
    // "iblog/router"
	"iblog/app"
) 

//入口文件
func main() {


    // r := gin.Default()
   
	// r.LoadHTMLGlob("templates/*")


	// router.RegistRouter(r)

	// r.Run(":8081") // listen and serve on 0.0.0.0:8080


    //启动应用 
	app.Run();


}