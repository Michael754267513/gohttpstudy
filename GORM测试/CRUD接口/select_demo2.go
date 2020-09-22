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
	var users []config.User
	db.Find(&users)
	for k, v := range users {
		fmt.Println(k, v)
	}

}
