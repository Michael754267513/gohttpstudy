package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Tag struct {
	ID     uint   `gorm:"primaryKey"`
	Locale string `gorm:"primaryKey"`
	Value  string
}
type Blog struct {
	ID         uint   `gorm:"primaryKey"`
	Locale     string `gorm:"primaryKey"`
	Subject    string
	Body       string
	Tags       []Tag `gorm:"many2many:blog_tags;"`
	SharedTags []Tag `gorm:"many2many:shared_blog_tags;ForeignKey:id;References:id"`
	LocaleTags []Tag `gorm:"many2many:locale_blog_tags;ForeignKey:id,locale;References:id"`
}

func main() {
	dsn := "root:laoshu@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Tag{}, &Blog{})
	//db.Create(&Blog{
	//	Locale:     "laoshu",
	//	Subject:    "wodigema",
	//	Body:       "body",
	//	Tags:       []Tag{
	//		{Value: "laoshu1"},
	//		{Value: "laoshu2"},
	//
	//	},
	//	SharedTags: []Tag{
	//		{Value: "SharedTags"},
	//		{Value: "SharedTags2"},
	//
	//	},
	//	LocaleTags: []Tag{
	//		{Value: "LocaleTags1"},
	//		{Value: "LocaleTags2"},
	//
	//	},
	//})
	//db.Create(&Blog{
	//	Locale:     "laoshu2",
	//	Subject:    "wodigema2",
	//	Body:       "body2",
	//	Tags:       []Tag{
	//		{Value: "laoshu11"},
	//		{Value: "laoshu22"},
	//
	//	},
	//	SharedTags: []Tag{
	//		{Value: "SharedTags11"},
	//		{Value: "SharedTags22"},
	//
	//	},
	//	LocaleTags: []Tag{
	//		{Value: "LocaleTags11"},
	//		{Value: "LocaleTags22"},
	//
	//	},
	//})
	var blog []Blog
	//db.Find(&blog,"body","body2")
	//db.Preload("SharedTags").Preload("LocaleTags").Preload("Tags").Find(&blog,"body","body")

	db.Preload(clause.Associations).Find(&blog, "body", "body")

	for _, v := range blog {
		jdog, _ := json.Marshal(v)
		fmt.Println(string(jdog))
	}
}
