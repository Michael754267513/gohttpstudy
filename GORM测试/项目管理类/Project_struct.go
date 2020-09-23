package 项目管理类

// 项目属性
type Project struct {
	Name          string
	ProductManger ProductManger
	ProjectManger ProjectManger
	TestManager   TestManager
	Developer     Developer
	Tester        Tester
	ProductLine   ProductLine
	SubLine       SubLine
	Desc          string
	Department    Department
	ProjectType   []ProjectType
}

// 部门名称
type Department struct {
	DepartmentName string
}

// 产品经理
type ProductManger struct {
	UserName []UserName
}

// 项目主事人员
type ProjectManger struct {
	UserName []UserName
}

// 测试主事人员
type TestManager struct {
	UserName []UserName
}

// 开发工程师
type Developer struct {
	UserName []UserName
}

// 测试工程师
type Tester struct {
	UserName []UserName
}

// 产品线(大方向)
type ProductLine struct {
	Name    string
	Product Product
}

// 产品子系统
type SubLine struct {
}

// 存放用户
type UserName struct {
	Name string
}

// 存在哪些产品
type Product struct {
}
