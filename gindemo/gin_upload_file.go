package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	// 默认32M
	router.MaxMultipartMemory = 8 << 20 // 8M
	router.POST("/upload_one", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.SaveUploadedFile(file, "sign_file")
		c.JSON(200, gin.H{
			"filename":   file.Filename,
			"fileHeader": file.Header,
		})
	})
	router.POST("/upload_multi", func(c *gin.Context) {
		//file,_ :=c.FormFile("file")
		//c.SaveUploadedFile(file,"sign_file")
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			fmt.Println("文件名: ", file)
			//c.SaveUploadedFile(file,dst)//存储文件
		}
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	router.Run()

	/*
		curl -X POST http://localhost:8080/upload \
		  -F "upload[]=@/Users/appleboy/test1.zip" \
		  -F "upload[]=@/Users/appleboy/test2.zip" \
		  -H "Content-Type: multipart/form-data"


	*/
}
