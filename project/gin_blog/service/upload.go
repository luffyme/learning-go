package service

import (
	"net/http"

	"gin_blog/pkg/e"
	"gin_blog/pkg/util"
	"gin_blog/pkg/setting"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	_, image, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : e.ERROR,
			"msg" : e.GetMsg(e.ERROR),
			"data" : nil,
		})
		return
	}

	if image == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code" : e.INVALID_PARAMS,
			"msg" : e.GetMsg(e.INVALID_PARAMS),
			"data" : nil,
		})
		return
	}

	imageName := util.GetImageName(image.Filename)
	fullPath := setting.Config.Upload.ImageUploadPath
	src := fullPath + imageName

	/*
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code" : e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT,
			"msg" : e.GetMsg(e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT),
			"data" : nil,
		})
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : e.ERROR_UPLOAD_CHECK_IMAGE_FAIL,
			"msg" : e.GetMsg(e.ERROR_UPLOAD_CHECK_IMAGE_FAIL),
			"data" : nil,
		})
		return
	}*/

	if err := c.SaveUploadedFile(image, src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : e.ERROR_UPLOAD_SAVE_IMAGE_FAIL,
			"msg" : e.GetMsg(e.ERROR_UPLOAD_SAVE_IMAGE_FAIL),
			"data" : nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : e.SUCCESS,
		"msg" : e.GetMsg(e.SUCCESS),
		"data" : map[string]string{
			//"image_url":      upload.GetImageFullUrl(imageName),
			"image_save_url": fullPath + imageName,
		},
	})
}