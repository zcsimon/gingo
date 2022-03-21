package main

import (
	"box/app/blog"
	"box/app/shop"
	"box/library"
	"box/router"
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

//读取key=value类型的配置文件
func InitConfig() map[string]string {
	config := make(map[string]string)

	f, err := os.Open(".env")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}
func main() {
	//

	config := InitConfig()

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
	fmt.Println(config)
	r.Run(":" + config["port"])
}
