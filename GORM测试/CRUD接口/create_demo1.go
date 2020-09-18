package main

import (
	"fmt"

	config "gohttpstudy/GORM测试"
)

// 默认会以ID 作为主键，如果结构体没有指明，会默认生成
// `gorm:"primaryKey"` 指定主键
func main() {
	db, err := config.MysqlClient()
	if err != nil {
		fmt.Println(err)
		panic("数据库连接错误：")
	}
	db.AutoMigrate(&config.User{})
}
