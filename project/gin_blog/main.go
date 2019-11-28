package main

import (
	"os"
	"io"
	"fmt"
	"time"
	_ "context"
	_ "os/signal"
	_ "net/http"
	_ "net/http/pprof"

	"gin_blog/models"
	"gin_blog/router"
	"gin_blog/middleware"

	_ "github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/fvbock/endless"
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

	router.Add(r)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/sleep", func(c *gin.Context) {
		fmt.Println("sleep fmt")

		time.Sleep(60 * time.Second)

		c.JSON(200, gin.H{
			"message": "sleep fmt",
		})
	})

	endless.ListenAndServe(":8000", r)

	/*
	router.Add(r)
	srv := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		//r.Run(":8000") // listen and serve on 0.0.0.0:8080
		srv.ListenAndServe()
	}()
	
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// quit 信道是同步信道，若没有信号进来，处于阻塞状态
	// 反之，则执行后续代码
	<-quit
	fmt.Println("start to Shutdown Server...");

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 调用 srv.Shutdown() 完成优雅停止
	// 调用时传递了一个上下文对象，对象中定义了超时时间
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err);
	}
	fmt.Println("Server Exit");
	*/
}
