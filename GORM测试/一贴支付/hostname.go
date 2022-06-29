/*
		Handpay ServiceMesh

           创建时间: 2020年11月25日15:55:24

	       少侠好武功,一起Giao起来
	  	 我说一Giao,你说Giao
		   一 Giao ？？？？

*/

package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Hosts struct {
	gorm.Model
	Namesapce string     `json:"namesapce"`
	Hostname  []Hostname `gorm:"foreignKey:HostID"`
}

type Hostname struct {
	gorm.Model
	Address string `json:"address"`
	Domain  string `json:"domain"`
	HostID  uint
}

func main() {
	dsn := "root:root@tcp(10.58.206.92:3306)/gva?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	//db.AutoMigrate(&Hosts{}, &Hostname{})

	//增
	//db.Create(&Hosts{
	//	Namesapce: "onepay-env2",
	//	Hostname: []Hostname{
	//		{Address: "10.148.181.201", Domain: "api.air.zr.common"},
	//	},
	//})

	// 查
	var hosts []Hosts
	//db.Preload(clause.Associations).Find(&hosts).Where("namespace = ?", "risk-air01")
	db.Preload(clause.Associations).Where("namespace = ?", "risk-air03").Find(&hosts)
	for _, v := range hosts {
		host, _ := json.Marshal(v)
		fmt.Println(string(host))
	}

	//hostname :=   Hostname{Address: "11.11.2.2",Domain: "ccc.abc.acom",HostID: 1}
	//db.Create(&hostname) //db.Delete(&hostname,2)
	//db.Delete(&hostname, []int{1,2,3})

	// 删除

}
