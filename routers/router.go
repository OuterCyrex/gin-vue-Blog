//这里用于配置router文件
//由于前后端分离项目需要大量的router
//我们需要单独配置

package routers

import (
	"fmt"
	"gin-vue-blog/api/version1"
	"gin-vue-blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	v1 := r.Group("api/version1")
	{
		//User模块的router接口
		v1.POST("user/add", version1.AddUser)
		v1.GET("users", version1.GetUser)
		v1.PUT("user/:id", version1.EditUser)
		v1.DELETE("user/:id", version1.DeleteUser)

		//Category模块的router接口

		v1.POST("category/add", version1.AddCategory)
		v1.GET("categories", version1.GetCategory)
		v1.PUT("category/:id", version1.EditCategory)
		v1.DELETE("category/:id", version1.DeleteCategory)

		//Article模块的router接口
		v1.POST("article/add", version1.AddArticle)
		v1.GET("articles", version1.GetArticle)
		v1.PUT("article/:id", version1.EditArticle)
		v1.DELETE("article/:id", version1.DeleteArticle)

	}
	err := r.Run(utils.HttpPort)
	fmt.Printf("routers出错，%v", err)
}
