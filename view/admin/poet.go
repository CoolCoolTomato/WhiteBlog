package admin

import (
	"WhiteBlog/common"
	"WhiteBlog/config"
	"WhiteBlog/models"
	"github.com/gin-gonic/gin"
)

// Poet 获取诗词
func Poet(c *gin.Context) {
	theConfig := config.GetConfig()
	poets, err := models.GetPoets()
	if err != nil {
		common.ServerErr(c, "Get Poets Fail")
		return
	}
	common.Render(c, "admin/poet.html", gin.H{
		"poets":  poets,
		"config": theConfig,
	})
}

// PoetAdd 新增诗词
func PoetAdd(c *gin.Context) {
	poet := models.Poet{}
	err := c.ShouldBindJSON(&poet)
	if err != nil || poet.Content == "" {
		common.BadRequest(c, "Data Error")
		return
	}
	err = poet.CreatePoet()
	if err != nil {
		common.ServerErr(c, "Create Poet Fail")
		return
	}
	common.Ok(c, "Create Success")
}

// PoetEdit 编辑诗词
func PoetEdit(c *gin.Context) {
	poet := models.Poet{}
	err := c.ShouldBindJSON(&poet)
	if err != nil || poet.Content == "" {
		common.BadRequest(c, "Data Error")
		return
	}
	err = poet.UpdatePoet()
	if err != nil {
		common.ServerErr(c, "Create Poet Fail")
		return
	}
	common.Ok(c, "Create Success")
}

// PoetDelete 删除诗词
func PoetDelete(c *gin.Context) {
	poet := models.Poet{}
	err := c.ShouldBindJSON(&poet)
	if err != nil {
		common.BadRequest(c, "Invalid Data")
		return
	}
	err = poet.DeletePoet()
	if err != nil {
		common.ServerErr(c, "Delete Fail")
		return
	}
	common.Ok(c, "Delete Success")
}
