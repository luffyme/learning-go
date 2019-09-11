package service

import (
	"strconv"
	"net/http"

	"gin_blog/models"
	"gin_blog/pkg/e"
	"gin_blog/pkg/util"

	"github.com/gin-gonic/gin"
)

type user struct{}
var User user

func (u *user) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.Output(c, http.StatusOK, e.ERROR_USER_NO_FOUND, nil)
		return
	}

	user, err := models.GetUser(id)
	if err != nil {
		util.Output(c, http.StatusOK, e.ERROR_USER_NO_FOUND, nil)
		return
	}

	util.Output(c, http.StatusOK, e.SUCCESS, user)
	return
}