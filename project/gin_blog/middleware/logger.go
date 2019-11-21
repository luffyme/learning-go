package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerToFile() gin.HandlerFunc {
	fileName := "logs/logger_access.log"

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()
	
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})


	return func(c *gin.Context) {
		startTime := time.Now()
		
		c.Next()
		
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		logger.WithFields(logrus.Fields{
			"statusCode": statusCode,
			"latencyTime": latencyTime,
			"clientIP": clientIP,
			"reqMethod": reqMethod,
			"reqUri": reqUri,
		  }).Info("A group of walrus emerges from the ocean")
		//logger.Infof("| %s | %s | %s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqUri)
		fmt.Println(statusCode, latencyTime, clientIP, reqMethod, reqUri)
	}
}

func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}