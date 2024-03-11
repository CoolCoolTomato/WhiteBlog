package models

import (
	"WhiteBlog/config"
	"time"
)

type Class struct {
	BaseModel
	Belong   int
	Name     string `gorm:"unique;not null"`
	Subclass []int  `gorm:"-"`
	//  分类包含的文章
	Articles []Article
}

type FormatClass struct {
	ID          int
	Name        string
	Belong      int
	Parent      string
	Subclasses  []Class
	SubclassNum int
	Articles    []string
	ArticleNum  int
	CreatedDate time.Time
	UpdatedDate time.Time
}

// GetClass 获取某个分类
func (class *Class) GetClass() error {
	return config.GetDatabase().First(&class).Error
}

// GetClasses 获取所有分类
func GetClasses() ([]FormatClass, error) {
	var classes []Class
	var formatClasses []FormatClass
	err := config.GetDatabase().Find(&classes).Error
	if err != nil {
		return formatClasses, err
	}
	for _, class := range classes {
		subclasses, _ := class.GetSubclasses()
		formatClass := FormatClass{
			ID:          class.BaseModel.ID,
			Name:        class.Name,
			Belong:      class.Belong,
			Parent:      "无",
			Subclasses:  subclasses,
			SubclassNum: len(subclasses),
			Articles:    nil,
			ArticleNum:  0,
			CreatedDate: class.BaseModel.CreatedDate,
			UpdatedDate: class.BaseModel.UpdatedDate,
		}
		formatClasses = append(formatClasses, formatClass)
	}
	return formatClasses, err
}

// GetRootClasses 获取一级分类
func GetRootClasses() ([]FormatClass, error) {
	var classes []Class
	var rootClasses []FormatClass
	err := config.GetDatabase().Where("belong = ?", -1).Find(&classes).Error
	if err != nil {
		return rootClasses, err
	}
	for _, class := range classes {
		subclasses, _ := class.GetSubclasses()
		articles, _ := GetArticlesByClass(class.BaseModel.ID)
		articleNum := len(articles)
		formatClass := FormatClass{
			ID:          class.BaseModel.ID,
			Name:        class.Name,
			Belong:      class.Belong,
			Parent:      "无",
			Subclasses:  subclasses,
			SubclassNum: len(subclasses),
			Articles:    nil,
			ArticleNum:  articleNum,
			CreatedDate: class.BaseModel.CreatedDate,
			UpdatedDate: class.BaseModel.UpdatedDate,
		}
		rootClasses = append(rootClasses, formatClass)
	}
	return rootClasses, err
}

// GetSubclasses 获取所有子分类
func (class *Class) GetSubclasses() ([]Class, error) {
	id := class.ID
	var subclasses []Class
	//  获取子分类
	err := config.GetDatabase().Where("belong = ?", id).Find(&subclasses).Error
	if err != nil {
		return subclasses, err
	}
	return subclasses, nil
}

// CreateClass 新增分类
func (class *Class) CreateClass() error {
	return config.GetDatabase().Create(&class).Error
}

// UpdateClass 更新某个分类
func (class *Class) UpdateClass() error {
	return config.GetDatabase().Updates(&class).Error
}

// DeleteClass 删除某个分类
func (class *Class) DeleteClass() error {
	return config.GetDatabase().Delete(&class).Error
}

// GetClassNum 获取分类数
func GetClassNum() int {
	var classes []Class
	err := config.GetDatabase().Find(&classes).Error
	if err != nil {
		return 0
	}
	return len(classes)
}
