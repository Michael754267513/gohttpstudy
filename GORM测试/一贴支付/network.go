/*
		Handpay ServiceMesh

           创建时间: 2020年11月25日15:55:24

	       少侠好武功,一起Giao起来
	  	 我说一Giao,你说Giao
		   一 Giao ？？？？

*/

package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type App struct {
	gorm.Model
	System        string          `json:"system"`
	SubSystem     string          `json:"sub_system"`
	NetworkPolicy []NetworkPolicy `gorm:"foreignKey:AppID"`
	Network       bool            `json:"network"`    // 是否开启网络权限
	StartTime     int             `json:"start_time"` // 启动时间单位秒
	Service       bool            `json:"service"`    // 是否开始service
	Replicas      int32           `json:"replicas"`   // 副本数
	CPU           int             `json:"cpu"`        // CPU
	Mem           int             `json:"mem"`        // 内存
	Probe         string          `json:"probe"`      // 检测url 检测返回http code
	Offline       string          `json:"offline"`    // 服务下线接口
	Skywalking    bool            `json:"skywalking"`
	Domain        []Domain        `gorm:"foreignKey:DomainID"`
}

type Domain struct {
	Domain      string `json:"domain"`
	ServiceName string `json:"service_name"`
	Port        int    `json:"port"`
	Uri         string `json:"uri"`
	DomainID    uint   `json:"domain_id"`
}

type DefaultAppPolicy struct {
	gorm.Model
	Env string `json:"env"` // k8s 集群环境

}

type DefaultNetworkPolicy struct {
	gorm.Model
	Name    string `json:"name"`    // 策略用途
	Address string `json:"address"` // ip地址
}

type NetworkPolicy struct {
	gorm.Model
	AppID       uint   `json:"app_id"`
	SrcSystem   string `json:"src_system"`
	DestSystem  string `json:"dest_system"`
	SrcAddress  string `json:"src_address"`
	DestAddress string `json:"dest_address"`
	Port        string `json:"port"`        // 端口
	PolicyType  string `json:"policy_type"` // 策略类型 Inbound Internel Outbound
}

type TestEnv struct {
	gorm.Model
	SubSystem string `json:"sub_system"`
	JavaOpts  string `json:"java_opts"`
}

//
//type Hosts struct {
//	gorm.Model
//	Namesapce 		string 			`json:"namesapce"`
//	Hostname		[]Hostname		`gorm:"foreignKey:HostID"`
//}
//
//type Hostname struct {
//	gorm.Model
//	Address 		string			`json:"address"`
//	Domain			string 			`json:"domain"`
//	HostID 			uint
//}

func main() {
	dsn := "root:laoshu@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&Hosts{},
		&Hostname{},
		&App{},
		&NetworkPolicy{},
		&Domain{},
		&TestEnv{},
		&DefaultAppPolicy{},
		&DefaultNetworkPolicy{},
	)

}
