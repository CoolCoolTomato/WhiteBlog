package admin

import (
	"WhiteBlog/common"
	"WhiteBlog/config"
	"WhiteBlog/models"
	"github.com/gin-gonic/gin"
)

// Link 友链
func Link(c *gin.Context) {
	theConfig := config.GetConfig()
	links, err := models.GetLinks()
	if err != nil {
		common.ServerErr(c, "Get Links Fail")
	}
	common.Render(c, "admin/link.html", gin.H{
		"links":  links,
		"config": theConfig,
	})
}

// LinkAdd 新增友链
func LinkAdd(c *gin.Context) {
	link := models.Link{}
	err := c.ShouldBindJSON(&link)
	if err != nil || link.Title == "" {
		common.ServerErr(c, "Data Error")
		return
	}
	err = link.CreateLink()
	if err != nil {
		common.ServerErr(c, "Create Link Fail")
		return
	}
	common.Ok(c, "Create Success")
}

// LinkEdit 编辑友链
func LinkEdit(c *gin.Context) {
	link := models.Link{}
	err := c.ShouldBindJSON(&link)
	if err != nil || link.Title == "" {
		common.ServerErr(c, "Data Error")
		return
	}
	err = link.UpdateLink()
	if err != nil {
		common.ServerErr(c, "Create Link Fail")
		return
	}
	common.Ok(c, "Create Success")
}

// LinkDelete 删除友链
func LinkDelete(c *gin.Context) {
	link := models.Link{}
	err := c.ShouldBindJSON(&link)
	if err != nil {
		common.BadRequest(c, "Invalid Data")
		return
	}
	err = link.DeleteLink()
	if err != nil {
		common.ServerErr(c, "Delete Fail")
		return
	}
	common.Ok(c, "Delete Success")
}
