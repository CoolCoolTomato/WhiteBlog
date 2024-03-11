package routes

import (
	"WhiteBlog/middleware"
	"WhiteBlog/view/admin"
	"WhiteBlog/view/article"
	"WhiteBlog/view/index"
	"WhiteBlog/view/link"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("sessions", store))
	//  登录
	r.GET("/login", admin.Login)
	r.POST("/login", admin.Login)
	r.GET("/logout", admin.Logout)
	//
	//  后台
	//
	adminRoute := r.Group("/admin")
	adminRoute.Use(middleware.AuthMiddleware())
	//  首页
	adminRoute.GET("/", admin.Index)
	adminRoute.GET("/index", admin.Index)
	//  分类
	{
		adminClass := adminRoute.Group("/class")
		adminClass.GET("/", admin.Class)
		adminClass.GET("/add", admin.ClassAdd)
		adminClass.POST("/add", admin.ClassAdd)
		adminClass.GET("/edit", admin.ClassEdit)
		adminClass.POST("/edit", admin.ClassEdit)
		adminClass.DELETE("/delete", admin.ClassDelete)
	}
	//  文章
	{
		adminArticle := adminRoute.Group("/article")
		adminArticle.GET("/", admin.Article)
		adminArticle.GET("/add", admin.ArticleAdd)
		adminArticle.POST("/add", admin.ArticleAdd)
		adminArticle.GET("/edit", admin.ArticleEdit)
		adminArticle.POST("/edit", admin.ArticleEdit)
		adminArticle.DELETE("/delete", admin.ArticleDelete)
		adminArticle.POST("/uploadimg", admin.UploadImg)
		adminArticle.DELETE("/deleteimg", admin.DeleteImg)
	}
	//  友链
	{
		adminLink := adminRoute.Group("/link")
		adminLink.GET("/", admin.Link)
		adminLink.POST("/add", admin.LinkAdd)
		adminLink.POST("/edit", admin.LinkEdit)
		adminLink.DELETE("/delete", admin.LinkDelete)
	}
	//  诗词
	{
		adminPoet := adminRoute.Group("/poet")
		adminPoet.GET("/", admin.Poet)
		adminPoet.POST("/add", admin.PoetAdd)
		adminPoet.POST("/edit", admin.PoetEdit)
		adminPoet.DELETE("/delete", admin.PoetDelete)
	}
	//  文件
	{
		adminFile := adminRoute.Group("/file")
		adminFile.GET("/", admin.File)
		adminFile.POST("/upload", admin.UploadFile)
		adminFile.DELETE("/delete", admin.DeleteFile)
	}
	//  设置
	adminRoute.GET("/setting", admin.Setting)
	adminRoute.POST("/setting", admin.Setting)
	//
	//  首页
	//
	r.GET("/", index.Index)
	r.GET("/index", index.Index)
	//
	//  文章
	//
	r.GET("/article", article.Index)
	r.GET("/article/:id", article.Article)
	//
	//  友链
	//
	r.GET("/link", link.Link)
}
