package main

import "log"
import "github.com/gin-gonic/gin"

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

// http://127.0.0.1:8080/testing?name=michael&address=20
func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	err := c.BindQuery(&person)
	if err != nil {
		log.Println("====== Only Bind Query String ======")
	}
	c.JSON(200, person)
}
