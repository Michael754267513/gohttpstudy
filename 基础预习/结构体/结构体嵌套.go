package main

import (
	"encoding/json"
	"fmt"
)

type Book1 struct {
	rall
	Name string
	age  int // 小写开头 私有字段
}
type rall struct {
	Type string
}

func main() {
	var b Book1
	b.Name = "可见性"
	b.age = 20
	b.rall.Type = "有点小意思"
	fmt.Println(b)
	jsonb, _ := json.Marshal(b)
	fmt.Printf("%v", string(jsonb))
	//{{有点小意思} 可见性 20}
	//{"Type":"有点小意思","Name":"可见性"}

}
