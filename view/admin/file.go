package admin

import (
	"WhiteBlog/common"
	"WhiteBlog/config"
	"WhiteBlog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func File(c *gin.Context) {
	theConfig := config.GetConfig()
	images, err := models.GetImages()
	if err != nil {
		common.ServerErr(c, "Get Images Fail")
		return
	}
	files, err := models.GetFiles()
	if err != nil {
		common.ServerErr(c, "Get Files Fail")
		return
	}
	common.Render(c, "admin/file.html", gin.H{
		"config": theConfig,
		"images": images,
		"files":  files,
	})
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		common.BadRequest(c, "Invalid File")
		return
	}
	filePath := "uploads/files/" + file.Filename
	urlPath := config.GetConfig().Path + filePath
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		common.ServerErr(c, "Save File Fail")
		return
	}
	fileM := models.File{
		FileName: file.Filename,
		FilePath: urlPath,
	}
	err = fileM.AddFile()
	if err != nil {
		common.ServerErr(c, "Save FileData Fail")
		return
	}
	c.JSON(http.StatusOK, fileM)
}

func DeleteFile(c *gin.Context) {
	file := models.File{}
	err := c.ShouldBindJSON(&file)
	if err != nil {
		common.BadRequest(c, "Invalid Data")
		return
	}
	err = file.GetFile()
	if err != nil {
		common.ServerErr(c, "Get File Fail")
		return
	}
	err = file.DeleteFile()
	if err != nil {
		common.ServerErr(c, "Delete Fail")
		return
	}
	filepath := "uploads/files/" + file.FileName
	err = os.Remove(filepath)
	if err != nil {
		common.ServerErr(c, "Delete Fail")
		return
	}
	common.Ok(c, "Delete Success")
}
