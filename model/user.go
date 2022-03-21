package model

import (
	"box/utils"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
)

type User struct {
	Id       int    `json:"id" form:"id" binding:"id"`
	UserName string `json:"username" form:"username" binding:"username"`
	Mobile   string `json:"mobile" form:"mobile" binding:"mobile"`
}

type UserRegister struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile" form:"mobile" binding:"mobile"`
	Email    string `json:"email"`
}

type LoginRequestParams struct {
	UserName string `form:"username" json:"username" uri:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
	//Ctx      *gin.Context
}

type SelectUserByIdParam struct {
	Id int `form:"id" json:"id" uri:"id" binding:"required"`
}

func CheckLogin(param *LoginRequestParams) (isPass bool, user *User, err error) {
	isPass = true
	// if param.UserName != "raw" {
	// 	isPass = false
	// 	err := errors.New("用户名或密码不正确!")
	// 	return false, User{}, err
	// }
	// return isPass, User{Id: 1, UserName: "zcSimon", Mobile: "13800138000"}, nil
	sqlString := "select id, username,password, mobile from tb_user where username = ?"

	// row, err := utils.Db.Prepare(sqlString)
	res := utils.Db.QueryRow(sqlString, param.UserName)

	var (
		id       int
		username string
		password string
		mobile   string
	)
	err1 := res.Scan(&id, &username, &password, &mobile)

	if err1 != nil {
		isPass = false
		return isPass, &User{}, errors.New("用户不存在")
	}

	h := md5.New()
	h.Write([]byte(param.Password))
	pmd5 := hex.EncodeToString(h.Sum(nil))
	log.Println(pmd5, password)
	if pmd5 != password {
		isPass = false
		return isPass, &User{}, errors.New("密码不正确")
	}
	return isPass, &User{Id: id, UserName: username, Mobile: mobile}, nil
}
