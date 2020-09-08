package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/query", func(c *gin.Context) {
		name := c.DefaultQuery("name", "haha") //设置查询默认值
		namespace := c.Query("namespace")
		c.JSON(200, gin.H{
			name:      name,
			namespace: namespace,
		})
	})
	router.POST("/query", func(c *gin.Context) {
		name := c.DefaultQuery("name", "haha") //设置查询默认值
		namespace := c.Query("namespace")
		c.JSON(200, gin.H{
			name:      name,
			namespace: namespace,
		})
	})
	router.Run()
}
