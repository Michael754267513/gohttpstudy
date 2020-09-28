package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get",
		})
	})
	router.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "post"})
	})
	//  很多很多

	router.Run()
}


