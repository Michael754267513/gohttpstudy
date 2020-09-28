package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "gohttpstudy/gindemo/swagger_demo/docs"
)

// @Tags User
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePassword true "用户修改密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /ping [PUT]
func main() {
	r := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/ping", pong)

	r.Run() // listen and serve on 0.0.0.0:8080
}

// Ping Pong
// @Summary Ping Pong
// @Description get string by ID
// @Tags 健康检测
// @Produce  json
// @Success 200
// @Router /ping [get]
func pong(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
