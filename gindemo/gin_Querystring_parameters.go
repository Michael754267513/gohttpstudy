package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/login", func(c *gin.Context) {
		name := c.DefaultQuery("name", "haha") //设置查询默认值
		namespace := c.Query("namespace")
		c.JSON(200, gin.H{
			name:      name,
			namespace: namespace,
		})
	})
	router.POST("/login", func(c *gin.Context) {
		name := c.DefaultQuery("name", "haha") //设置查询默认值
		namespace := c.Query("namespace")
		c.JSON(200, gin.H{
			name:      name,
			namespace: namespace,
		})
	})
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
	}()
	router.Run()
}
