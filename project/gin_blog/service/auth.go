package service

import (
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
)

type Auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func (a *Auth) GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	atuhVa := a{Username: username, Password: password}
	ok, _ := valid.Valid(&atuhVa)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err = util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
	})
}