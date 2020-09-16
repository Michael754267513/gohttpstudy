package main

import (
	"fmt"
	"regexp"

	"github.com/asaskevich/govalidator"
)

// https://zhuanlan.zhihu.com/p/222307590
//需要自定义自定义的structTag,类似json:"id"
// govalidator
type User struct {
	Id       int    `valid:"type(int)"`
	Name     string `valid:"in(hgha|gggg|hzeng),stringlength(2|10)"` // 验证字段
	Name1    string `valid:"-"`                                      // - 不验证
	Bio      string `valid:"email,optional"`                         // 字段可以为空，如果不为空需要验证email
	Email    string `valid:"email"`                                  //
	Message  string `valid:"duck,ascii"`                             // 自定义
	Message2 string `valid:"animal(dog)"`                            // 带参数
	TstMail  string `valid:"required~email必须是754267513@qq.com"`      //自定义错误返回

}

func main() {
	var p1 User
	p1.Id = 200
	p1.Name = "hzeng"
	p1.Email = "754267513@qq.com"
	p1.Message2 = "dog"
	p1.Message = "duck"
	// 自定义验证
	govalidator.TagMap["duck"] = govalidator.Validator(func(str string) bool {
		return str == "duck" //判断值是否是duck
	})
	// 自定义带参数的验证
	govalidator.ParamTagMap["animal"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		species := params[0]  // 获取传送的参数
		return str == species // 判断是都等于str dog的
	})
	govalidator.ParamTagRegexMap["animal"] = regexp.MustCompile("^animal\\((\\w+)\\)$")

	result, err := govalidator.ValidateStruct(p1)
	fmt.Println(p1)
	if err != nil {
		println("error: " + err.Error())
	}
	println(result)
	fmt.Println(p1)
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}
