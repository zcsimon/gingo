package middle

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("mi", "middleware")
		c.Next()
		status := c.Writer.Status()
		log.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		log.Println("time:", t2)
	}
}
