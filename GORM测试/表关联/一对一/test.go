package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type One struct {
	gorm.Model
	Name string
	Two  Two `gorm:"foreignkey:TwoID;references:id"`
}

type Two struct {
	gorm.Model
	TwoName  string
	UserName []Username `gorm:"foreignkey:UsernameID;references:id"`
	TwoID    int
}

type Username struct {
	ListName   string
	UsernameID int
}

func main() {
	dsn := "root:laoshu@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&One{},
		&Two{},
		&Username{},
	)
	//
	//db.Create(&One{
	//	Name: "haha",
	//	Two:  Two{
	//		TwoName:  "TwoName",
	//		UserName: []Username{
	//			{ListName: "mama"},
	//			{ListName: "baba"},
	//		},
	//	},
	//})
	//db.Create(&One{
	//	Name: "haha1",
	//	Two:  Two{
	//		TwoName:  "TwoName1",
	//		UserName: []Username{
	//			{ListName: "mama1"},
	//			{ListName: "baba1"},
	//		},
	//	},
	//})
	var one []One
	////db.Preload(clause.Associations).Find(one)
	////for _,v := range one {
	////	jdog,_ := json.Marshal(v)
	////	fmt.Println(string(jdog))
	////}
	//
	//db.Preload("two").Find(&one)
	//fmt.Println(one)
	db.Preload("Two.UserName").Preload("Two").Find(&one, "name", "haha1")
	//db.Debug().Preload("Two.UserName").Preload("Two").Find(&one,"name","haha1")
	//db.Raw("").Scan(&one)
	for _, v := range one {
		jdog, _ := json.Marshal(v)
		fmt.Println(string(jdog))
	}

}
