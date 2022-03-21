package dao

import (
	"box/model"
	"box/utils"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
)

func CheckLogin(param *model.LoginRequestParams) (isPass bool, user *model.User, err error) {
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
		return isPass, &model.User{}, errors.New("用户不存在")
	}

	h := md5.New()
	h.Write([]byte(param.Password))
	pmd5 := hex.EncodeToString(h.Sum(nil))
	log.Println(pmd5, password)
	if pmd5 != password {
		isPass = false
		return isPass, &model.User{}, errors.New("密码不正确")
	}
	return isPass, &model.User{Id: id, UserName: username, Mobile: mobile}, nil
}
