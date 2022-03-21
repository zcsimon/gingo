package model

import (
	"errors"
)

type User struct {
	Id       int    `json:"id" form:"id" binding:"id"`
	UserName string `json:"username" form:"username" binding:"username"`
	Mobile   string `json:"mobile" form:"mobile" binding:"mobile"`
}

type LoginRequestParams struct {
	UserName string `form:"username" json:"username" uri:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
	//Ctx      *gin.Context
}

type SelectUserByIdParam struct {
	Id int `form:"id" json:"id" uri:"id" binding:"required"`
}

func CheckLogin(param *LoginRequestParams) (isPass bool, user User, err error) {
	isPass = true
	if param.UserName != "raw" {
		isPass = false
		err := errors.New("用户名或密码不正确!")
		return false, User{}, err
	}
	return isPass, User{Id: 1, UserName: "zcSimon", Mobile: "13800138000"}, nil

}
