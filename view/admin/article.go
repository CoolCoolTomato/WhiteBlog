package admin

import (
	"WhiteBlog/common"
	"WhiteBlog/config"
	"WhiteBlog/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Form struct {
	Article models.Article
	Image   []int
}

// Article 文章
func Article(c *gin.Context) {
	theConfig := config.GetConfig()
	classes, err := models.GetClasses()
	if err != nil {
		common.ServerErr(c, "Get Classes Fail")
		return
	}
	articles, err := models.GetArticles()
	if err != nil {
		common.ServerErr(c, "Get Articles Fail")
		return
	}
	classId, _ := strconv.Atoi(c.Query("class"))
	class := models.Class{}
	class.ID = classId
	err = class.GetClass()
	if err == nil && classId > 0 {
		articles, err = models.GetArticlesByClass(class.ID)
		if err != nil {
			common.ServerErr(c, "Get Articles Fail")
			return
		}
	} else {
		class = models.Class{}
	}
	common.Render(c, "admin/article.html", gin.H{
		"config":   theConfig,
		"class":    class,
		"classes":  classes,
		"articles": articles,
	})
}

// ArticleAdd 新增文章
func ArticleAdd(c *gin.Context) {
	theConfig := config.GetConfig()
	//  post请求
	if c.Request.Method == "POST" {
		form := Form{}
		err := c.ShouldBindJSON(&form)
		if err != nil {
			common.BadRequest(c, "Data Error")
			return
		}
		err = form.Article.CreateArticle()
		if err != nil {
			common.ServerErr(c, "Create Article Fail")
			return
		}
		models.UpdateArticleID(form.Image, form.Article.ID)
		common.Ok(c, "Add Success")
		return
	}
	// get请求
	classes, err := models.GetRootClasses()
	if err != nil {
		common.ServerErr(c, "Get Classes Fail")
		return
	}
	common.Render(c, "admin/articleAdd.html", gin.H{
		"classes": classes,
		"config":  theConfig,
	})
}

// ArticleEdit 编辑文章
func ArticleEdit(c *gin.Context) {
	theConfig := config.GetConfig()
	//  post请求
	if c.Request.Method == "POST" {
		form := Form{}
		err := c.ShouldBindJSON(&form)
		if err != nil || form.Article.Title == "" {
			common.BadRequest(c, "Data Error")
			return
		}
		err = form.Article.UpdateArticle()
		if err != nil {
			common.ServerErr(c, "Update Class Err")
			return
		}
		models.UpdateArticleID(form.Image, form.Article.ID)
		common.Ok(c, "Update Success")
		return
	}
	//  get请求
	article := models.Article{}
	articleId, _ := strconv.Atoi(c.Query("article"))
	if articleId <= 0 {
		common.BadRequest(c, "Invalid ID")
		return
	}
	article.ID = articleId
	err := article.GetArticle()
	if err != nil {
		common.ServerErr(c, "Get Article Fail")
		return
	}
	classes, err := models.GetRootClasses()
	if err != nil {
		common.ServerErr(c, "Get Classes Fail")
		return
	}
	images, err := models.GetImagesByArticleID(articleId)
	if err != nil {
		common.ServerErr(c, "Get Images Fail")
		return
	}
	common.Render(c, "admin/articleEdit.html", gin.H{
		"article": article,
		"classes": classes,
		"images":  images,
		"config":  theConfig,
	})
}

// ArticleDelete 删除文章
func ArticleDelete(c *gin.Context) {
	article := models.Article{}
	err := c.ShouldBindJSON(&article)
	if err != nil {
		common.BadRequest(c, "Invalid Data")
		return
	}
	err = article.DeleteArticle()
	if err != nil {
		common.ServerErr(c, "Delete Fail")
		return
	}
	common.Ok(c, "Delete Success")
}
