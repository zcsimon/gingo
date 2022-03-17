package v1

import (
	"box/library"
	"box/services"
	"log"

	"github.com/gin-gonic/gin"
)

// json 数据接收及绑定 shouldBind  or BindJson
func Login(c *gin.Context) {
	responseBody := library.NewResponseBody()
	defer library.RecoverResponse(c, responseBody)
	//c.PostForm("text")
	param := &services.LoginRequestParams{}
	//fmt.Println(param)
	// text = c.PostForm("text")
	//err := c.BindJSON(param)
	err := c.ShouldBind(param)
	if err != nil {
		log.Println(err.Error())
		responseBody.Code = 500
		responseBody.Message = err.Error()
		return
	}
	services.UserLogin(param, responseBody)

}

// Form 表单绑定
func LoginFromForm(c *gin.Context) {
	responseBody := library.NewResponseBody()
	defer library.RecoverResponse(c, responseBody)
	//c.PostForm("text")
	param := &services.LoginRequestParams{}
	//fmt.Println(param)
	// text = c.PostForm("text")
	//err := c.BindJSON(param)
	err := c.Bind(param)
	if err != nil {
		log.Println(err.Error())
		responseBody.Code = 500
		responseBody.Message = err.Error()
		return
	}
	log.Printf("%v", param.UserName)
	services.UserLoginFromForm(param, responseBody)

}

// Uri 参数绑定。
func LoginFromUri(c *gin.Context) {
	responseBody := library.NewResponseBody()
	defer library.RecoverResponse(c, responseBody)
	//c.PostForm("text")
	param := &services.LoginRequestParams{}
	//fmt.Println(param)
	// text = c.PostForm("text")
	//err := c.BindJSON(param)
	err := c.ShouldBindUri(param)
	if err != nil {
		log.Println(err.Error())
		responseBody.Code = 500
		responseBody.Message = err.Error()
		return
	}
	services.UserLoginFromForm(param, responseBody)

}

// TODO:: uri参数 query参数及 Form 参数同时请求 处理方案。
