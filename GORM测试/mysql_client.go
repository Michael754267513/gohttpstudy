package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlClient() (db *gorm.DB, err error) {

	dsn := "root:laoshu@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}

// 单表测试
type User struct {
	gorm.Model
	Name string `gorm:"primaryKey"` // 设置主键
	// 0、''、false 之类零值，这些字段定义的默认值不会被保存到数据库，您需要使用指针类型或 Scanner/Valuer 来避免这个问题
	Age int `gorm:"default:18"` // 设置默认值
}
