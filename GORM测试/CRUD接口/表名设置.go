package main

import (
	"fmt"

	config "gohttpstudy/GORM测试"
	"gorm.io/gorm"
)

// 默认会以ID 作为主键，如果结构体没有指明，会默认生成
// 默认使用结构体的复数作为表名 比如TableName 表名则为TableNames
type TableName struct {
	gorm.Model
	Name string
}

// 设置表名
func (TableName) TableName() string {
	return "test_tableName"
}
func main() {
	db, err := config.MysqlClient()
	if err != nil {
		fmt.Println(err)
		panic("数据库连接错误：")
	}
	db.AutoMigrate(&TableName{})
}
