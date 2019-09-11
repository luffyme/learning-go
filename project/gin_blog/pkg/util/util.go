package util

import (
	"path"
	"strings"

	"gin_blog/pkg/e"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = EncodeMD5(fileName)

	return fileName + ext
}

func Output(c *gin.Context, httpCode int, errCode int, data interface{}) {
	c.JSON(httpCode, Result{
		Code : errCode,
		Msg : e.GetMsg(errCode),
		Data : data,
	})
	return
}