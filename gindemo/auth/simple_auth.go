package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(AuthMiddleWare())
	{

		r.GET("/login", func(c *gin.Context) {
			cookie := &http.Cookie{
				Name:     "session_id",
				Value:    "onion",
				Path:     "/",
				HttpOnly: true,
			}
			http.SetCookie(c.Writer, cookie)
			c.String(http.StatusOK, "登录成功")
		})

		r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello world"})
		})
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL.String())
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			fmt.Println(value)
			if value == "onion" {
				c.Next()
				return
			}
		}
		if url := c.Request.URL.String(); url == "/login" {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
}
