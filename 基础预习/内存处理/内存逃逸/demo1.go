package main

// 编译参数-gcflag=-m可以查看编译过程中的逃逸分析：

type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) //局部变量s逃逸到堆

	s.Name = name
	s.Age = age

	return s
}

func main() {
	StudentRegister("Jim", 18)
}

/*
C:\Users\hzeng\go\src\gohttpstudy\基础预习\内存处理\内存逃逸>go build -gcflags=-m demo1.go
# command-line-arguments
.\demo1.go:10:6: can inline StudentRegister
.\demo1.go:19:6: can inline main
.\demo1.go:20:17: inlining call to StudentRegister
.\demo1.go:10:22: leaking param: name
.\demo1.go:11:10: new(Student) escapes to heap // 逃逸到堆
.\demo1.go:20:17: new(Student) does not escape
*/
