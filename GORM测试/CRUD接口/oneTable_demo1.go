package main

import (
	"database/sql"
	"fmt"

	config "gohttpstudy/GORM测试"
)

type Project struct {
	ID     int
	Name   string
	Build  string
	Age    int          `gorm:"default:18"` // 默认值
	Active sql.NullBool `gorm:"default:true"`
}

func main() {
	var pstringList []Project
	db, err := config.MysqlClient()
	if err != nil {
		fmt.Println(err)
		panic("数据库连接错误：")
	}
	db.AutoMigrate(&Project{})

	// 插入数据
	//var p Project
	//p.Name = "t3"
	//p.Build = "b3"
	//db.Save(&p)

	// 结构体插入
	//p2 := Project{
	//	Name:  "t4",
	//	Build: "t4",
	//}
	//db.Save(&p2)

	// slice 插入
	//p3 := []Project{{Name: "t5",Build: "b5"},{Name: "t6",Build: "b6"}}
	//db.Create(&p3)

	// map 插入
	//p4 := map[string]interface{}{"name":"map","build":"build_map"}
	//db.Model(&Project{}).Create(&p4)

	// slice 类型的map
	//p5 := []map[string]interface{}{
	//	{"name":"map1","build":"build_ma1p"},
	//	{"name":"map2","build":"build_map2"},
	//}
	//db.Model(&Project{}).Create(&p5)

	//查询
	//var pselct []Project
	//db.Find(&pselct) // 查询所有 ，返回的是数组
	//fmt.Println(pselct)

	// 条件查询
	//var ps1 []Project
	//db.Find(&ps1,[]int{1,2,3})  // SELECT * FROM projects WHERE id IN (1,2,3);
	//fmt.Println(ps1[0].Active.Bool)

	//  where 字符串条件查询
	//db.Where("name = ?","t1").First(&pstringList)
	//db.Find(&pstringList,"name = ?","t1")

	//db.Where("name like ?","map").Find(&pstringList)
	//db.Find(&pstringList,"name like ?","map")

	//db.Where("name in ?",[]string{"t1","t2"}).Find(&pstringList)
	//db.Find(&pstringList,"name in ?",[]string{"t1","t2"})

	// map slice 查询
	//db.Where(&Project{Name: "t2"}).Find(&pstringList)
	//db.Where(map[string]interface{}{"name": "t1"}).Find(&pstringList)
	//db.Where([]int{2,3}).Find(&pstringList)

	// not 条件
	//db.Not("name = ?","t1").Find(&pstringList)
	//db.Not("name in ?",[]string{"t1","t2"}).Find(&pstringList)

	// or 条件
	//db.Where("name = ?","t1").Or("name = ?","t2").Find(&pstringList)
	//db.Where(map[string]interface{}{"name": "t1"}).Or(&Project{Name: "t2"}).Find(&pstringList)

	//选择特定字段
	//db.Select("name").Find(&pstringList)
	// SELECT name,   FROM projects;

	// order
	//db.Order("name desc").Find(&pstringList)
	//db.Order("name desc").Order("age").Find(&pstringList)
	//SELECT * FROM projects ORDER BY name desc, age;

	//limit limit 限制 offset偏移量， offset >0 向后便宜，offset <0 向前便宜
	//db.Limit(3).Find(&pstringList)
	//db.Limit(3).Offset(-1).Find(&pstringList)

	// group  having

	// Distinct 唯一值查询
	//db.Distinct("age","name").Find(&pstringList)

	// joins

	fmt.Println(pstringList)
}
