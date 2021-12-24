package app

import (
	"ApiForCode/app/common/request"
	"ApiForCode/app/common/response"
	"ApiForCode/app/services"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Info(c *gin.Context) {
	var form request.Code
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	temp, err := services.GetCode(form)
	if err == nil {
		response.Success(c, temp)
	} else {
		response.BusinessFail(c, err.Error())
	}
}

func Img2Base64(c *gin.Context) {
	f, err := c.FormFile("img")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	img, err := f.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, img); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	temp, err := services.GetBase64(buf.Bytes())
	if err == nil {
		response.Success(c, temp)
	} else {
		response.BusinessFail(c, err.Error())
	}
}
