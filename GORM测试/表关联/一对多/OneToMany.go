package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Dog struct {
	ID   int
	Name string
	Toy  []Toy `gorm:"polymorphic:Owner;"`
}
type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

func main() {
	dsn := "root:laoshu@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var dog []Dog

	db.AutoMigrate(&Dog{}, &Toy{})
	// 迁移 schema
	//db.Create(&Dog{Name: "dog2", Toy: []Toy{{Name: "toy11"}, {Name: "toy22"}}})
	//db.Find(&dog)
	//db.Preload("Toy").Find(&dog)
	db.Preload(clause.Associations).Find(&dog)
	//jdog,err := json.Marshal(dog)
	//fmt.Println(string(jdog))
	for _, v := range dog {
		jdog, _ := json.Marshal(v)
		fmt.Println(string(jdog))
	}
}
