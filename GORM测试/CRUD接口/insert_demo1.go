package main

import (
	"fmt"

	config "gohttpstudy/GORM测试"
)

// GORM 允许 BeforeSave, BeforeCreate, AfterSave, AfterCreate 等钩子
//func (u config.User) BeforeCreate(tx *gorm.DB) (err error) {
//
//
//	if u.Role == "admin" {
//		return errors.New("invalid role")
//	}
//	return
//}
func main() {
	db, err := config.MysqlClient()
	if err != nil {
		fmt.Println(err)
		panic("数据库连接错误：")
	}
	var user config.User
	user = config.User{
		Name: "hzeng",
		Age:  20,
	}
	res := db.Create(&user) // 必须用指针
	fmt.Printf("返回插入的状态：%v", res.Error)
	fmt.Printf("\n返回插入的行数：%v", res.RowsAffected)
	// 选定字段创建
	user_select := config.User{
		Name: "hzeng1",
		Age:  18,
	}
	db.Select("Name", user_select)
	// 批量插入
	var users = []config.User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)

	// GORM 支持根据 map[string]interface{} 和 []map[string]interface{}{} 创建记录
	userMap := []map[string]interface{}{
		{"Name": "michael", "Age": 20},
		{"Name": "zenghao", "Age": 10},
	}
	db.Model(&config.User{}).Create(&userMap)
}
