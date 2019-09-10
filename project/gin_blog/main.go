package main

import (
	"gin_blog/router"
	
	"github.com/gin-gonic/gin"
   	_ "github.com/davecgh/go-spew/spew"
) 

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