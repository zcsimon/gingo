package main

import (
	"box/app/blog"
	"box/app/shop"
	"box/library"
	"box/router"
	"box/utils"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//

	// 初始化配置文件
	env := utils.InitConfig()

	gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("runtime/logs/gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 1.创建路由 。由 Include 及 Init 替代
	//r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello World!")
	// })
	// 注册不同目录\app、模块下的路由
	router.Include(shop.Router, blog.Router, router.Http)
	// 初始化
	r := router.Init()
	//router.Http(r)
	//r.Use(middle.Middleware())

	r.NoRoute(func(c *gin.Context) {
		resBody := &library.ResponseBody{Code: 404, Message: "route not found!!!"}
		c.JSON(http.StatusNotFound, resBody)
	})

	// 3.启动服务，监听端口
	r.Run(":" + env["port"])

}
