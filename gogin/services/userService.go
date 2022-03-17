package services

// 服务层 。TODO:: gorm操作
import (
	"box/library"
)

type LoginRequestParams struct {
	UserName string `form:"username" json:"username" uri:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
	//Ctx      *gin.Context
}

func UserLogin(param *LoginRequestParams, responseBody *library.ResponseBody) {
	// TODO::Gorm 操作及jwt
	responseBody.SetData(map[string]string{"username": param.UserName, "from": "china hebei shijiazhuang"})

	return
}

func UserLoginFromForm(param *LoginRequestParams, responseBody *library.ResponseBody) {
	// TODO::Gorm 操作及jwt
	responseBody.SetData(map[string]string{"username": param.UserName, "from": "china hebei shijiazhuang"})

	return
}

func UserLoginFromUri(param *LoginRequestParams, responseBody *library.ResponseBody) {
	// TODO::Gorm 操作及jwt
	responseBody.SetData(map[string]string{"username": param.UserName, "from": "china hebei shijiazhuang"})

	return
}
