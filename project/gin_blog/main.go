package main

import (
	"os"
	"io"

	"gin_blog/models"
	"gin_blog/router"
	"gin_blog/middleware"

	_ "github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

func init() {
	models.Setup()
}

func main() {
	//将日志记录在文件中
	f, _ := os.Create("logs/access.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.Use(middleware.LoggerToFile())

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Add(r)

	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
