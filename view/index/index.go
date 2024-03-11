package index

import (
	"WhiteBlog/common"
	"WhiteBlog/config"
	"WhiteBlog/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	theConfig := config.GetConfig()
	poets, err := models.GetPoets()
	if err != nil {
		common.ServerErr(c, "Get Poet Fail")
	}
	common.Render(c, "index/index.html", gin.H{
		"config": theConfig,
		"poets":  poets,
	})
}
