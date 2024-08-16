package version1

import (
	"gin-vue-blog/model"
	"gin-vue-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加友链

func AddFriendLink(c *gin.Context) {
	var data model.FriendLink
	_ = c.ShouldBindJSON(&data)
	code := model.CreateFriendLink(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询友链列表

func GetFriendLink(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	//Cancel limit and offset if value == -1 是gorm的特性

	data, code, total := model.GetFriendLink(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除友链

func DeleteFriendLink(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteFriendLink(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
