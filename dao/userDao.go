package dao

import (
	"box/model"
	"box/utils"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

func CheckLogin(param *model.LoginRequestParams) (isPass bool, user *model.User, err error) {
	isPass = true
	// method 1
	sqlString := "select id, username,password, mobile from tb_user where username = ?"

	//row, err := utils.Db.Prepare(sqlString)
	row := utils.DbConn.QueryRow(sqlString, param.UserName)
	userInfo := &model.UserRegister{}
	err = row.Scan(&userInfo.ID, &userInfo.Username, &userInfo.Password, &userInfo.Mobile)

	if err != nil {
		isPass = false
		return isPass, &model.User{}, errors.New("用户不存在")
	}
	// method 2
	// mysql := &utils.Mysql{}
	// data := make(map[string]interface{})

	// mysql.Sql = "select id, username,password, mobile from tb_user where username = " + string("\""+param.UserName+"\"")
	// rows, err := utils.Db.Query(mysql)
	// if len(rows) == 0 && err == nil {
	// 	isPass = false
	// 	return isPass, &model.User{}, errors.New("用户不存在")
	// }
	// userInfo := &model.UserRegister{}
	// for _, row := range rows {
	// 	userInfo.ID, _ = strconv.Atoi(row["id"])
	// 	userInfo.Username = row["username"]
	// 	userInfo.Mobile = row["mobile"]
	// 	userInfo.Password = row["password"]

	// }
	//err1 := row.Scan(&userInfo.ID, &userInfo.Username, &userInfo.Password, &userInfo.Mobile)

	// if err1 != nil {
	// 	isPass = false
	// 	return isPass, &model.User{}, errors.New("用户不存在")
	// }

	h := md5.New()
	h.Write([]byte(param.Password))
	pmd5 := hex.EncodeToString(h.Sum(nil))
	if pmd5 != userInfo.Password {
		isPass = false
		return isPass, &model.User{}, errors.New("密码不正确")
	}
	return isPass, &model.User{Id: userInfo.ID, UserName: userInfo.Username, Mobile: userInfo.Mobile}, nil
}
