package main

import (
	"gin_blog/models"
	"gin_blog/router"

	_ "github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

func init() {
	models.Setup()
}

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Add(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
