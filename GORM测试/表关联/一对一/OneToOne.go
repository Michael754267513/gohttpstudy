package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard `gorm:"foreignKey:CreditCardID",ASSOCIATION_FOREIGNKEY:ID"`
	// 使用 UserName 作为外键
}
type CreditCard struct {
	gorm.Model
	Number       string
	UserName     string
	CreditCardID int
}

func main() {
	dsn := "root:laoshu@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &CreditCard{})
	//db.Create(&User{Name: "hzeng", CreditCard: CreditCard{UserName: "toy11",Number: "nu12"}})
	//db.Create(&User{Name: "michael", CreditCard: CreditCard{UserName: "toymichael",Number: "numichael"}})
	var user []User

	//db.Find(&user,"name","hzeng").Preload("CreditCard")
	db.Preload(clause.Associations).Find(&user)

	for _, v := range user {
		jdog, _ := json.Marshal(v)
		fmt.Println(string(jdog))
	}

}
