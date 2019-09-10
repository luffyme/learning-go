package router

import (
	"gin_blog/service"

	"github.com/gin-gonic/gin"
)

func Add(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/auth", service.Auth.GetAuth)

	return r
}