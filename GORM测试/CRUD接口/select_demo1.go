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

	if res := db.First(&config.User{}); res != nil {
		fmt.Printf("db.First: %v", res)
		// SELECT * FROM users ORDER BY id LIMIT 1;
	}
	if res := db.Take(&config.User{}); res != nil {
		fmt.Printf("\ndb.Take: %v", res)
		// SELECT * FROM users LIMIT 1;
	}
	if res := db.Last(&config.User{}); res != nil {
		fmt.Printf("\ndb.Last: %v", res)
		// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	}
	if res := db.Find(&config.User{}); res != nil {
		fmt.Printf("\ndb.Find: %v", res)
		// SELECT * FROM users ;
	}

	// 条件查找
	if res := db.First(&config.User{}, 2); res != nil {
		fmt.Printf("\n条件where: %v", res)
		//  SELECT * FROM users WHERE id = 2;
	}
	if res := db.First(&config.User{}, []int{1, 2, 4}); res != nil {
		fmt.Printf("\n条件where: %v", res)
		//  SELECT * FROM users WHERE id IN (1,2,4);
	}
	//// 获取第一条匹配的记录
	//db.Where("name = ?", "jinzhu").First(&user)
	//// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	//
	//// 获取全部匹配的记录
	//db.Where("name <> ?", "jinzhu").Find(&users)
	//// SELECT * FROM users WHERE name <> 'jinzhu';
	//
	//// IN
	//db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');
	//
	//// LIKE
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	//// SELECT * FROM users WHERE name LIKE '%jin%';
	//
	//// AND
	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	//// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	//
	//// Time
	//db.Where("updated_at > ?", lastWeek).Find(&users)
	//// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
	//
	//// BETWEEN
	//db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	////SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	// 结构体和map

	//// Struct
	//db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
	//
	//// Map
	//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
	//
	//// 主键切片条件
	//db.Where([]int64{20, 21, 22}).Find(&users)
	//// SELECT * FROM users WHERE id IN (20, 21, 22);

	//db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
}
