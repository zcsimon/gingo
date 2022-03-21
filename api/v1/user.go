package v1

import (
	"box/library"
	"box/model"
	"box/services"
	"log"

	"github.com/gin-gonic/gin"
)

var identityKey = "id"

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

// json 数据接收及绑定 shouldBind  or BindJson
func Login(c *gin.Context) {
	responseBody := library.NewResponseBody()
	defer library.RecoverResponse(c, responseBody)
	//c.PostForm("text")
	param := &model.LoginRequestParams{}
	//fmt.Println(param)
	// text = c.PostForm("text")
	//err := c.BindJSON(param)
	err := c.ShouldBind(param)
	if err != nil {
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
	param := &model.LoginRequestParams{}
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
	param := &model.LoginRequestParams{}
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

func GetUserById(c *gin.Context) {
	// claims := c.MustGet("claims").(*jwt.CustomClaims)
	// if claims != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": 0,
	// 		"msg":    "token有效",
	// 		"data":   claims,
	// 	})
	// }
	log.Println("start controller")
	responseBody := library.NewResponseBody()

	defer library.RecoverResponse(c, responseBody)
	//c.PostForm("text")
	param := &model.SelectUserByIdParam{}

	err := c.BindUri(param)
	log.Println("controller")
	if err != nil {
		log.Println(err.Error())
		responseBody.Code = 500
		responseBody.Message = err.Error()
		return
	}

	services.GetUserById(c, param, responseBody)

}
