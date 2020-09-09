package handler

import (
	"github.com/gin-gonic/gin"
	. "gohttpstudy/gindemo/gin_casbin/component"
)

func ReadResource(c *gin.Context) {
	// some stuff
	// blahblah...

	c.JSON(200, RestResponse{Code: 1, Message: "read resource successfully", Data: "resource"})
}

func WriteResource(c *gin.Context) {
	// some stuff
	// blahblah...
	write := c.PostForm("write")
	c.JSON(200, RestResponse{Code: 1, Message: "write resource successfully", Data: write})
}
