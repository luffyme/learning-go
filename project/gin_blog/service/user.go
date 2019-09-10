package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct{}
var User user

func (u *user) GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name" : "luffyme",
		"age" : 18,
	})
}