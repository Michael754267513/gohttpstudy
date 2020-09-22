package main

import (
	"fmt"

	config "gohttpstudy/GORM测试"
)

func main() {
	db, err := config.MysqlClient()
	if err != nil {
		fmt.Println(err)
		panic("数据库连接错误：")
	}
	var user config.User
	db.First(&user)

}
