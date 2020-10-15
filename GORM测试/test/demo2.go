package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Project struct {
	gorm.Model
	Name string `json:"NameID" form:"NameID" gorm:"column:name;comment:项目名称;type:varchar(256);size:256;"`
	Desc string `json:"Desc" form:"Desc" gorm:"column:desc;comment:描述;type:text(1024);size:1024;"`
	//ProductManger ProductManger `json:"ProductManger" form:"ProductManger" gorm:"column:productmanger;comment:产品经理;foreignKey:ProductMangerID,ASSOCIATION_FOREIGNKEY:ID"`
	//ProjectManger ProjectManger `json:"ProjectManger" form:"ProjectManger" gorm:"column:projectmanger;comment:项目经理;foreignKey:ProjectMangerID,ASSOCIATION_FOREIGNKEY:ID"`
	//TestManager   TestManager	`json:"TestManager" form:"TestManager" gorm:"column:testmanager;comment:测试经理;foreignKey:TestManagerID,ASSOCIATION_FOREIGNKEY:ID"`
	//Developer     Developer		`json:"Developer" form:"Developer" gorm:"column:developer;comment:项目经理;foreignKey:DeveloperID,ASSOCIATION_FOREIGNKEY:ID"`
	//Tester        Tester		`json:"Tester" form:"Tester" gorm:"column:tester;comment:项目经理;foreignKey:TesterID,ASSOCIATION_FOREIGNKEY:ID"`
	//ProductLine   ProductLine
	//SubLine       SubLine
	//Department    Department	`json:"Department" form:"Department" gorm:"column:department;comment:部门;foreignKey:DepartmentID,ASSOCIATION_FOREIGNKEY:ID"`
	Department Department `json:"Department" form:"Department" gorm:"column:department;comment:部门;foreignKey:DepartmentID"`
}

func (Project) TableName() string {
	return "project"
}

// 部门名称
type Department struct {
	gorm.Model
	DepartmentID   int    `json:"DepartmentID" form:"DepartmentID" gorm:"column:departmentID;comment:id;"`
	DepartmentName string `json:"DepartmentName" form:"DepartmentName" gorm:"column:DepartmentName;comment:项目名称;type:varchar(256);size:256;"`
}

// 产品经理
type ProductManger struct {
	gorm.Model
	UserName        []UserName `json:"UserName" form:"UserName" gorm:"column:username;comment:用户;foreignKey:ProductMangerID,ASSOCIATION_FOREIGNKEY:ID"`
	ProductMangerID int        `json:"ProductMangerID" form:"ProductMangerID" gorm:"column:ProductMangerID;comment:id;"`
}

// 项目主事人员
type ProjectManger struct {
	gorm.Model
	ProjectMangerID int        `json:"ProjectMangerID" form:"ProjectMangerID" gorm:"column:ProjectMangerID;comment:id;"`
	UserName        []UserName `json:"UserName" form:"UserName" gorm:"column:username;comment:用户;foreignKey:ProjectMangerID,ASSOCIATION_FOREIGNKEY:ID"`
}

// 测试主事人员
type TestManager struct {
	gorm.Model
	TestManagerID int        `json:"TestManagerID" form:"TestManagerID" gorm:"column:TestManagerID;comment:id;"`
	UserName      []UserName `json:"UserName" form:"UserName" gorm:"column:username;comment:用户;foreignKey:TestManagerID,ASSOCIATION_FOREIGNKEY:ID"`
}

// 开发工程师
type Developer struct {
	gorm.Model
	DeveloperID int        `json:"DeveloperID" form:"DeveloperID" gorm:"column:DeveloperID;comment:id;"`
	UserName    []UserName `json:"UserName" form:"UserName" gorm:"column:username;comment:用户;foreignKey:DeveloperID,ASSOCIATION_FOREIGNKEY:ID"`
}

// 测试工程师
type Tester struct {
	gorm.Model
	TesterID int        `json:"TesterID" form:"TesterID" gorm:"column:TesterID;comment:id;"`
	UserName []UserName `json:"UserName" form:"UserName" gorm:"column:username;comment:用户;foreignKey:TesterID,ASSOCIATION_FOREIGNKEY:ID"`
}

// 产品线(大方向)
type ProductLine struct {
	gorm.Model
	ProductLineID int
	Name          string
	Product       Product
}

// 产品子系统
type SubLine struct {
}

// 存放用户
type UserName struct {
	//gorm.Model
	Name            string `json:"Name" form:"Name" gorm:"column:Name;comment:用户名;"`
	ProjectMangerID int    `json:"ProjectMangerID" form:"ProjectMangerID" gorm:"column:ProjectMangerID;comment:id;"`
	ProductMangerID int    `json:"ProductMangerID" form:"ProductMangerID" gorm:"column:ProductMangerID;comment:id;"`
	DeveloperID     int    `json:"DeveloperID" form:"DeveloperID" gorm:"column:DeveloperID;comment:id;"`
	TestManagerID   int    `json:"TestManagerID" form:"TestManagerID" gorm:"column:TestManagerID;comment:id;"`
	TesterID        int    `json:"TesterID" form:"TesterID" gorm:"column:TesterID;comment:id;"`
}

// 存在哪些产品
type Product struct {
}

func main() {
	dsn := "root:laoshu@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&Project{},
		//&ProductManger{},
		//&ProjectManger{},
		//&TestManager{},
		//&Tester{},
		//&Developer{},
		&Department{},
	)
}
