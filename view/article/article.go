package article

import (
	"WhiteBlog/common"
	"WhiteBlog/config"
	"WhiteBlog/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Index(c *gin.Context) {
	theConfig := config.GetConfig()
	classes, err := models.GetRootClasses()
	if err != nil {
		common.ServerErr(c, "Get Classes Fail")
	}
	var articles []models.FormatArticle
	var class models.Class
	classId, _ := strconv.Atoi(c.Query("class"))
	if classId > 0 {
		class.ID = classId
		err = class.GetClass()
	} else {
		err = nil
	}
	if err != nil || classId <= 0 {
		articles, err = models.GetArticles()
		if err != nil {
			common.ServerErr(c, "Get Articles Fail")
			return
		}
	} else {
		articles, err = models.GetArticlesByClass(classId)
		if err != nil {
			common.ServerErr(c, "Get Articles Fail")
			return
		}
	}
	common.Render(c, "article/index.html", gin.H{
		"config":   theConfig,
		"classes":  classes,
		"class":    class.Name,
		"articles": articles,
	})
}

func Article(c *gin.Context) {
	theConfig := config.GetConfig()
	id, _ := strconv.Atoi(c.Param("id"))
	article := models.Article{}
	article.ID = id
	err := article.GetArticle()
	if err != nil {
		common.ServerErr(c, "Get Article Fail")
		return
	}
	common.Render(c, "article/article.html", gin.H{
		"config":  theConfig,
		"article": article,
	})
}
