package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	g1 := router.Group("/g1")
	{
		g1.GET("/get", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Get"})
		})
		g1.POST("/post", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Post"})
		})
	}
	g2 := router.Group("/g2")
	{
		g2.GET("/get", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "g2Get"})
		})
		g2.POST("/post", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "g2Post"})
		})
	}
	router.Run()
}
