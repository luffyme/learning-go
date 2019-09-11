package router

import (
	"gin_blog/service"
	"gin_blog/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func Add(r *gin.Engine) *gin.Engine {
	r.GET("/ping", service.Ping)
	r.GET("/auth", service.GetAuth)
	r.GET("/upload", service.UploadImage)
	r.GET("/user/:id", jwt.JWT(), service.User.GetUser)

	return r
}