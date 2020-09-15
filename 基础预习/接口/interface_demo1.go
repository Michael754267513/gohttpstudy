package main

import "fmt"

// 接口使用场景，支付方式有多种，用同一个接口去实现
type Action interface {
	Say(s string) (err error)
}

type Git struct {
}
type Svn struct {
}

func (action Git) Say(s string) (err error) {
	fmt.Println(s)
	return err
}

func (action Svn) Say(s string) (err error) {
	fmt.Println(s)
	return err
}

func main() {
	git := Git{}
	svn := Svn{}
	var A Action
	A = git
	A.Say("git")
	A = svn
	A.Say("svn")

}
