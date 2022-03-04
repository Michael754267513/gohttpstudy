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
	dsn := "root:laoshu@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Hosts{}, &Hostname{})

	//增
	db.Create(&Hosts{
		Namesapce: "risk-air03",
		Hostname: []Hostname{
			{Address: "10.148.181.219", Domain: "crypto-host-a"},
			{Address: "10.148.181.219", Domain: "crypto-host-b"},
			{Address: "110.148.181.204", Domain: "api.air.gateway.zr.common"},
			{Address: "110.148.181.201", Domain: "api.air.engine.zr.common"},
			{Address: "10.148.181.201", Domain: "admin.air.engine.zr.common"},
			{Address: "10.148.181.201", Domain: "api.air.norm.zr.common"},
			{Address: "10.148.181.201", Domain: "admin.air.norm.zr.common"},
			{Address: "10.148.181.203", Domain: "api.1tpayinside.zr.common"},
		},
	})

	// 查
	var hosts []Hosts
	db.Preload(clause.Associations).Find(&hosts, 1)
	for _, v := range hosts {
		host, _ := json.Marshal(v)
		fmt.Println(string(host))
	}

	//hostname :=   Hostname{Address: "11.11.2.2",Domain: "ccc.abc.acom",HostID: 1}
	//db.Create(&hostname) //db.Delete(&hostname,2)
	//db.Delete(&hostname, []int{1,2,3})

	// 删除

}
