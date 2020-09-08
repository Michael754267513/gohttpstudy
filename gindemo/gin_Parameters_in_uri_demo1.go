package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/apis/:namespace/:name", func(c *gin.Context) {
		name := c.Param("name")
		namespace := c.Param("namespace")
		c.JSON(200, gin.H{
			name:      name,
			namespace: namespace,
		})
	}) //{"aaaa":"aaaa","k8s":"k8s"}
	router.GET("/api/:namespace/*name", func(c *gin.Context) {
		name := c.Param("name")
		namespace := c.Param("namespace")
		c.JSON(200, gin.H{
			name:      name,
			namespace: namespace,
		})
	}) // {"/aaaa":"/aaaa","k8s":"k8s"}
	router.Run()
}
