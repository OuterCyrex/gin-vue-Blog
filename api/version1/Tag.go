package version1

import (
	"gin-vue-blog/model"
	"gin-vue-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加标签

func AddTag(c *gin.Context) {
	var data model.Tag
	_ = c.ShouldBindJSON(&data)
	code := model.CheckTag(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateTag(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个标签

func SearchTag(c *gin.Context) {
	var data model.Tag
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	data, code := model.SearchTag(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询标签列表

func GetTag(c *gin.Context) {
	data, code, total := model.GetTag()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//通过分类获取标签

func GetTagByCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, total := model.GetTagByCategory(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑标签

func EditTag(c *gin.Context) {
	var data model.Tag
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.CheckUpTag(data.ID, data.Name)
	if code == errmsg.SUCCESS {
		code = model.EditTag(id, &data)
	}
	if code == errmsg.ERROR_TAG_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除分类

func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteTag(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
