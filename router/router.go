package router

import (
	v1 "box/api/v1"
	"box/library"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册路由
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	r := gin.New()
	for _, opt := range options {
		opt(r)
	}
	return r
}

// TODO:: 可拆分到其他目录路由
func Http(r *gin.Engine) {

	apiRouter := r.Group("api/v1")
	{
		apiRouter.POST("/login_json", v1.Login)
		apiRouter.POST("/login_form", v1.LoginFromForm)
		apiRouter.GET("/login/:username/:password", v1.LoginFromUri)
	}
	r.NoRoute(func(c *gin.Context) {
		resBody := &library.ResponseBody{Code: 404, Message: "route not found!"}
		c.JSON(http.StatusNotFound, resBody)
	})
}
