package services

// 服务层 。TODO:: gorm操作
import (
	"box/library"
	"box/middle"
	"box/model"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// type LoginRequestParams struct {
// 	UserName string `form:"username" json:"username" uri:"username" binding:"required"`
// 	Password string `form:"password" json:"password" uri:"password" binding:"required"`
// 	//Ctx      *gin.Context
// }

// LoginResult 登录结果结构
type LoginResult struct {
	Token   string `json:"token"`
	Expired int    `json:"expired"`
	model.User
}

var identityKey = "id"

func UserLogin(param *model.LoginRequestParams, responseBody *library.ResponseBody) {
	// TODO::Gorm 操作及jwt
	//responseBody.SetData(map[string]string{"username": param.UserName, "from": "china hebei shijiazhuang"})
	j := &middle.JWT{
		[]byte("zcSimon"),
	}

	isPass, user, err := model.CheckLogin(param)
	log.Println(isPass, user, err)
	if !isPass && err != nil {
		responseBody.SetCode(-1)
		responseBody.SetMessage(err.Error())
		return
	}

	//user := model.User{}
	claims := middle.CustomClaims{
		user.Id,
		user.UserName,
		user.Mobile,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "zcSimon",                       //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		responseBody.SetCode(-1)
		responseBody.SetMessage("创建token 失败")

		return
	}

	data := LoginResult{
		User:    user,
		Token:   token,
		Expired: int(claims.ExpiresAt),
	}

	responseBody.SetCode(http.StatusOK)
	responseBody.SetMessage("登录成功")
	responseBody.SetData(data)

	//responseBody.SetCode()
	return
}
func UserLoginFromForm(param *model.LoginRequestParams, responseBody *library.ResponseBody) {
	// TODO::Gorm 操作及jwt
	responseBody.SetData(map[string]string{"username": param.UserName, "from": "china hebei shijiazhuang"})

	return
}

func UserLoginFromUri(param *model.LoginRequestParams, responseBody *library.ResponseBody) {
	// TODO::Gorm 操作及jwt
	responseBody.SetData(map[string]string{"username": param.UserName, "from": "china hebei shijiazhuang"})

	return
}

func GetUserById(c *gin.Context, param *model.SelectUserByIdParam, responseBody *library.ResponseBody) {
	claims := c.MustGet("claims").(*middle.CustomClaims)
	if param.Id != claims.Id {
		responseBody.SetCode(-1)
		responseBody.SetMessage("参数不合法,请确定要查询数据的正确性!")
		return
	}
	if claims != nil {

		responseBody.SetCode(http.StatusOK)
		responseBody.SetData(claims)
		responseBody.SetMessage("查询用户数据成功")
	}
	return
}
