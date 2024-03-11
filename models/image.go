package models

import "WhiteBlog/config"

type Image struct {
	ID        int `gorm:"primary_key"`
	FileName  string
	FilePath  string
	ArticleID int
}
type FormatImage struct {
	ID           int
	FileName     string
	FilePath     string
	ArticleTitle string
}

func GetImages() ([]FormatImage, error) {
	var images []Image
	var formatImages []FormatImage
	err := config.GetDatabase().Find(&images).Error
	for _, image := range images {
		var articleTitle string
		article := Article{}
		article.ID = image.ArticleID
		err := article.GetArticle()
		if err != nil {
			articleTitle = "无"
		} else {
			articleTitle = article.Title
		}
		formatImage := FormatImage{
			ID:           image.ID,
			FileName:     image.FileName,
			FilePath:     image.FilePath,
			ArticleTitle: articleTitle,
		}
		formatImages = append(formatImages, formatImage)
	}
	return formatImages, err
}

// GetImage 查找图片
func (image *Image) GetImage() error {
	return config.GetDatabase().First(&image).Error
}

// AddImage 添加图片
func (image *Image) AddImage() error {
	return config.GetDatabase().Create(&image).Error
}

// UpdateImage 更新图片
func (image *Image) UpdateImage() error {
	return config.GetDatabase().Updates(&image).Error
}

// DeleteImage 删除图片
func (image *Image) DeleteImage() error {
	return config.GetDatabase().Delete(&image).Error
}

// UpdateArticleID 批量更新ArticleID
func UpdateArticleID(idList []int, articleID int) {
	for _, id := range idList {
		img := Image{}
		img.ID = id
		err := img.GetImage()
		if err != nil {
			continue
		}
		img.ArticleID = articleID
		err = img.UpdateImage()
		if err != nil {
			continue
		}
	}
}

// GetImagesByArticleID 根据文章ID获取图片
func GetImagesByArticleID(id int) ([]Image, error) {
	var images []Image
	err := config.GetDatabase().Where("article_id = ?", id).Find(&images).Error
	return images, err
}
