//这里用于配置router文件
//由于前后端分离项目需要大量的router
//我们需要单独配置

package routers

import (
	"fmt"
	"gin-vue-blog/api/version1"
	"gin-vue-blog/middleware"
	"gin-vue-blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//User模块的router接口

		auth.PUT("user/:id", version1.EditUser)
		auth.DELETE("user/:id", version1.DeleteUser)

		//Category模块的router接口

		auth.POST("category/add", version1.AddCategory)

		auth.PUT("category/:id", version1.EditCategory)
		auth.DELETE("category/:id", version1.DeleteCategory)

		//Article模块的router接口
		auth.POST("article/add", version1.AddArticle)

		auth.PUT("article/:id", version1.EditArticle)
		auth.DELETE("article/:id", version1.DeleteArticle)

		//Upload模块的router接口
		auth.POST("upload", version1.UploadFile)
	}
	router := r.Group("api/v1")
	{
		router.GET("users", version1.GetUser)
		router.GET("user/:id", version1.GetUserInfo)
		router.GET("categories", version1.GetCategory)
		router.GET("articles", version1.GetArticle)
		router.GET("article/:id", version1.SearchArticle)
		router.GET("article/category/:id", version1.GetArticleByCategory)
		router.POST("user/add", version1.AddUser)
		router.POST("login", version1.Login)
	}
	err := r.Run(utils.HttpPort)
	fmt.Printf("routers出错，%v", err)
}
