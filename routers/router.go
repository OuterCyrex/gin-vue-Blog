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
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.LoadHTMLFiles("static/admin/admin.html", "static/index.html")
	r.Static("admin/static", "static/admin/static")
	r.GET("admin/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", nil)
	})
	r.Static("static", "static/static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//User模块的router接口

		auth.PUT("user/:id", version1.EditUser)
		auth.DELETE("user/:id", version1.DeleteUser)
		auth.PUT("user/pwd/:id", version1.EditPWD)

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

		//Profile模块的router接口
		auth.PUT("profile/:id", version1.UpdateProfile)

		// FriendLink模块的router接口
		auth.POST("friend/add", version1.AddFriendLink)
		auth.DELETE("friend/:id", version1.DeleteFriendLink)

		//Comments模块的router接口
		auth.DELETE("comment/:id", version1.DeleteComment)
	}

	router := r.Group("api/v1")
	{
		//用户接口
		router.GET("users", version1.GetUser)
		router.GET("user/:id", version1.GetUserInfo)
		router.POST("user/add", version1.AddUser)
		router.GET("profile/:id", version1.GetProfile)
		//分类接口
		router.GET("categories", version1.GetCategory)
		router.GET("category/:id", version1.SearchCategory)
		//文章接口
		router.GET("articles", version1.GetArticle)
		router.GET("article/:id", version1.SearchArticle)
		router.GET("article/category/:id", version1.GetArticleByCategory)

		//评论接口
		router.GET("comment/:id", version1.GetCommentsByArticle)
		router.POST("comment", version1.AddComment)
		router.GET("comment", version1.GetComments)

		//登录接口
		router.POST("userlogin", version1.UserLogin)
		router.POST("login", version1.Login)
		router.GET("userinfo", version1.GetLoginInfoByToken)

		//友链接口
		router.GET("friend", version1.GetFriendLink)
	}
	err := r.Run(utils.HttpPort)
	fmt.Printf("routers出错，%v", err)
}
