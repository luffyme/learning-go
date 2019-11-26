package router

import (
	"gin_blog/service"
	"gin_blog/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/cors"
)

func Add(r *gin.Engine) *gin.Engine {
	pprof.Register(r, "debug/pprof")

	/*
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com"}
	r.Use(cors.New(config))
	*/

	//跨域
	r.Use(cors.Default())


	r.GET("/ping", service.Ping)
	r.GET("/auth", service.GetAuth)
	r.GET("/upload", service.UploadImage)
	r.GET("/user/:id", middleware.JWT(), service.User.GetUser)

	return r
}