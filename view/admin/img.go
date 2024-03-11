package admin

import (
	"WhiteBlog/common"
	"WhiteBlog/config"
	"WhiteBlog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func UploadImg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		common.BadRequest(c, "Invalid File")
		return
	}
	imgPath := "uploads/article/" + file.Filename
	urlPath := config.GetConfig().Path + imgPath
	err = c.SaveUploadedFile(file, imgPath)
	if err != nil {
		common.ServerErr(c, "Save Image Fail")
		return
	}
	image := models.Image{
		FileName:  file.Filename,
		FilePath:  urlPath,
		ArticleID: 0,
	}
	err = image.AddImage()
	if err != nil {
		common.ServerErr(c, "Save ImageData Fail")
		return
	}
	c.JSON(http.StatusOK, image)
}

func DeleteImg(c *gin.Context) {
	image := models.Image{}
	err := c.ShouldBindJSON(&image)
	if err != nil {
		common.BadRequest(c, "Invalid Data")
		return
	}
	err = image.GetImage()
	if err != nil {
		common.ServerErr(c, "Get Image Fail")
		return
	}
	err = image.DeleteImage()
	if err != nil {
		common.ServerErr(c, "Delete Fail")
		return
	}
	filepath := "uploads/article/" + image.FileName
	err = os.Remove(filepath)
	if err != nil {
		common.ServerErr(c, "Delete Fail")
		return
	}
	common.Ok(c, "Delete Success")
}
