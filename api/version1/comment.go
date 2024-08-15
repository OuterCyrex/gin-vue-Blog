package version1

import (
	"gin-vue-blog/model"
	"gin-vue-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCommentsByArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	aid, _ := strconv.Atoi(c.Param("id"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code, total := model.GetCommentsByArticle(pageSize, pageNum, aid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func AddComment(c *gin.Context) {
	var user model.User
	code, user := GetNameByToken(c)
	var comment model.Comments
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"data":    comment,
		})
	}
	_ = c.ShouldBind(&comment)
	comment.Uid = int(user.ID)
	code = model.CreateComment(&comment)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    comment,
		"message": errmsg.GetErrMsg(code),
	})
}
