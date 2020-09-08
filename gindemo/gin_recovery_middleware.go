package main

import "github.com/gin-gonic/gin"

func main() {
	//router := gin.Default() // default 默认是没有中间件的
	router := gin.New()
	router.Use(gin.Logger())   // 使用中间件
	router.Use(gin.Recovery()) // 发生panic 返回500
	//router.Use(gin.)
	router.GET("/panic", func(c *gin.Context) {
		panic("err")
	})
	router.Run()
}
