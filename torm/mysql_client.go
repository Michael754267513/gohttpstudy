package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func MysqlClient() (db *gorm.DB, err error) {

	dsn := "root:laoshu@tcp(127.0.0.1:3306)/orm?charset=utf8mb4&parseTime=True&loc=Local"
	// 重写logger
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second,   // Slow SQL threshold
	//		LogLevel:      logger.Silent, // Log level
	//		Colorful:      false,         // Disable color
	//	},
	//)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
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
