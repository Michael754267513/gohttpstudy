package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PC struct {
	gorm.Model
	Name string
	PB   PB `gorm:"foreignKey:PCID",ASSOCIATION_FOREIGNKEY:ID"`
}

type PB struct {
	gorm.Model
	Name string
	PCID int
	PA   []PA `gorm:"foreignKey:PAID",ASSOCIATION_FOREIGNKEY:ID"`
}

type PA struct {
	Name string
	PAID int
}

func main() {
	dsn := "root:laoshu@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&PC{}, &PB{}, &PA{})

}
