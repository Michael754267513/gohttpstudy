package 项目管理类

// 新建项目类型 代码工程 配置文件 静态文件 Sql文件
type ProjectType struct {
	Code Code
	//Config 			Config
	//StaticFile 		StaticFile
	//SQLFile			SQLFile
}

// 代码类型
type Code struct {
	CVSAddress  string
	CodeType    CodeType
	Branch      string
	PkgType     PkgType
	BuildType   BuildType
	BuildOrder  int  // 编译顺序
	IsDependent bool // 是否是依赖基础编译包
	PkgName     string
}

// svn git 或者其他版本控制器
type CodeType struct {
	CVS CVS
}

// springboot tomcat jboss
type PkgType struct {
	Pkg string
}

// 编译使用编译器  maven or ？
type BuildType struct {
	Build string
}

//// 配置文件
//type Config struct {
//
//}
//
//// 静态文件
//type StaticFile struct {
//
//}
//
//// sql文件
//type SQLFile struct {
//
//}
