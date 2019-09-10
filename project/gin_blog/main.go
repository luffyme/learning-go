package main

import (
	"gin_blog/router"

	"github.com/gin-gonic/gin"
   	_ "github.com/davecgh/go-spew/spew"
) 

func main() {
	r := gin.Default()
	router.Add(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}