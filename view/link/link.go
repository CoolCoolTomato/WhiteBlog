package link

import (
	"WhiteBlog/common"
	"WhiteBlog/config"
	"WhiteBlog/models"
	"github.com/gin-gonic/gin"
)

func Link(c *gin.Context) {
	theConfig := config.GetConfig()
	links, err := models.GetLinks()
	if err != nil {
		common.ServerErr(c, "Get Links Fail")
		return
	}
	common.Render(c, "link/link.html", gin.H{
		"config": theConfig,
		"links":  links,
	})
}
