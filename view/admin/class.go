package admin

import (
	"WhiteBlog/common"
	"WhiteBlog/config"
	"WhiteBlog/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Class 获取分类
func Class(c *gin.Context) {
	theConfig := config.GetConfig()
	classes, err := models.GetRootClasses()
	if err != nil {
		common.ServerErr(c, "Get Classes Fail")
		return
	}
	common.Render(c, "admin/class.html", gin.H{
		"classes": classes,
		"config":  theConfig,
	})
}

// ClassAdd 新增分类
func ClassAdd(c *gin.Context) {
	theConfig := config.GetConfig()
	if c.Request.Method == "POST" {
		class := models.Class{}
		err := c.ShouldBindJSON(&class)
		if err != nil || class.Name == "" {
			common.BadRequest(c, "Data Error")
			return
		}
		err = class.CreateClass()
		if err != nil {
			common.ServerErr(c, "Create Class Fail")
			return
		}
		common.Ok(c, "Create Success")
		return
	}

	class := models.Class{}
	classId, _ := strconv.Atoi(c.Query("id"))
	if classId <= 0 {
		class.ID = 0
	} else {
		class.ID = classId
		err := class.GetClass()
		if err != nil {
			class.ID = 0
		}
	}

	classes, err := models.GetRootClasses()
	if err != nil {
		common.ServerErr(c, "Get Classes Fail")
		return
	}
	subclasses, err := class.GetSubclasses()
	if err != nil {
		common.ServerErr(c, "Get Subclass Fail")
		return
	}
	common.Render(c, "admin/classAdd.html", gin.H{
		"class":      class,
		"classes":    classes,
		"subclasses": subclasses,
		"config":     theConfig,
	})
}

// ClassEdit 编辑分类
func ClassEdit(c *gin.Context) {
	theConfig := config.GetConfig()
	//  post请求
	if c.Request.Method == "POST" {
		class := models.Class{}
		err := c.ShouldBindJSON(&class)
		if err != nil || class.Name == "" {
			common.BadRequest(c, "Data Error")
			return
		}
		err = class.UpdateClass()
		if err != nil {
			common.ServerErr(c, "Update Class Err")
			return
		}
		common.Ok(c, "Update Success")
		return
	}
	//  get请求
	class := models.Class{}
	classId, _ := strconv.Atoi(c.Query("id"))
	if classId <= 0 {
		common.BadRequest(c, "Invalid ID")
		return
	}
	class.ID = classId
	err := class.GetClass()
	if err != nil {
		common.ServerErr(c, "Get Class Fail")
		return
	}

	classes, err := models.GetRootClasses()
	if err != nil {
		common.ServerErr(c, "Get Class Fail")
		return
	}
	subclasses, err := class.GetSubclasses()
	if err != nil {
		common.ServerErr(c, "Get Subclass Fail")
		return
	}
	common.Render(c, "admin/classEdit.html", gin.H{
		"class":      class,
		"classes":    classes,
		"subclasses": subclasses,
		"config":     theConfig,
	})
}

// ClassDelete 删除分类
func ClassDelete(c *gin.Context) {
	class := models.Class{}
	err := c.ShouldBindJSON(&class)
	if err != nil {
		common.BadRequest(c, "Invalid Data")
		return
	}
	err = class.DeleteClass()
	if err != nil {
		common.ServerErr(c, "Delete Fail")
		return
	}
	common.Ok(c, "Delete Success")
}
