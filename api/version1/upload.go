package version1

import (
	"gin-vue-blog/Server"
	"gin-vue-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, header, _ := c.Request.FormFile("file")
	fileSize := header.Size
	url, code := Server.UploadFile(file, fileSize)
	if code != errmsg.SUCCESS {
		c.JSON(200, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"url":  "",
		})
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"url":  url,
	})
}
