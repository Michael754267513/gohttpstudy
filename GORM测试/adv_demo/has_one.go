package main

import (
	"fmt"

	config "gohttpstudy/GORM测试"
)

type Cat struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
	Age       int
}

func main() {
	db, err := config.MysqlClient()
	if err != nil {
		fmt.Println(err)
		panic("数据库连接错误：")
	}
	db.AutoMigrate(&Cat{})
	db.AutoMigrate(&Dog{})
	db.AutoMigrate(&Toy{})
	//var cat Cat
	//cat.Toy.Name = "cat"
	//db.Save(&cat)
	//db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})
	// 数据插入
	//var test HasUser
	//test.CreditCard.UserName= "michael"
	//test.CreditCard.Number = "10"
	//db.Save(&test)
	//var test1 HasUser
	//db.Find(&test1)
	//fmt.Println(test1.CreditCard.UserName)

	var cat Dog
	db.First(&cat)
	fmt.Println(cat.Toy)
}
