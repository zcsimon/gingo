package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	// Db, err := sql.Open("mysql", "root:123456@localhost:3306/beego")
	Db, err = sql.Open("mysql", "root:123456@/beego")

	if err != nil {
		log.Println("连接失败:", err.Error())
		panic(err.Error())
	}

}
