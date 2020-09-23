package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type One struct {
	Name string
	Two  Two `gorm:"-"`
}

type Two struct {
	TwoName  string
	UserName []Username `gorm:"-"`
}

type Username struct {
	ListName string
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

	db.Create(&One{
		Name: "haha",
		Two: Two{
			TwoName: "TwoName",
			UserName: []Username{
				{ListName: "mama"},
				{ListName: "baba"},
			},
		},
	})
	db.Create(&One{
		Name: "haha1",
		Two: Two{
			TwoName: "TwoName1",
			UserName: []Username{
				{ListName: "mama1"},
				{ListName: "baba1"},
			},
		},
	})
	var one []One
	db.Preload(clause.Associations).Find(one)
	for _, v := range one {
		jdog, _ := json.Marshal(v)
		fmt.Println(string(jdog))
	}
}
