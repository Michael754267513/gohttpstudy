package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/query", func(c *gin.Context) {
		name := c.PostForm("name")
		namespce := c.PostForm("namespace")
		namedefault := c.DefaultPostForm("namedefault", "默认值") // 默认值
		nameMap := c.PostFormMap("nameMap")
		//nameArray,bool := c.GetPostFormArray("nameArray")
		c.JSON(200, gin.H{
			name:        name,
			namespce:    namespce,
			namedefault: namedefault,
			"nameMap":   nameMap,
		})
	})
	router.Run()
}
