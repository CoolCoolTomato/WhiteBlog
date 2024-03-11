package models

import (
	"WhiteBlog/config"
	"sort"
	"time"
)

type Article struct {
	BaseModel
	Title   string
	Content string
	//  外键
	ClassID int
}

type FormatArticle struct {
	ID          int
	Title       string
	Class       string
	CreatedDate time.Time
	UpdatedDate time.Time
}

// GetArticles 获取所有文章标题
func GetArticles() ([]FormatArticle, error) {
	var articles []Article
	var formatArticles []FormatArticle
	err := config.GetDatabase().Select([]string{"ID", "title", "ClassID", "CreatedDate", "UpdatedDate"}).Find(&articles).Error
	for _, article := range articles {
		class := Class{}
		class.ID = article.ClassID
		err := class.GetClass()
		if err != nil {
			continue
		}
		formatArticle := FormatArticle{
			ID:          article.ID,
			Title:       article.Title,
			Class:       class.Name,
			CreatedDate: article.CreatedDate,
			UpdatedDate: article.UpdatedDate,
		}
		formatArticles = append(formatArticles, formatArticle)
	}
	for i, j := 0, len(formatArticles)-1; i < j; i, j = i+1, j-1 {
		formatArticles[i], formatArticles[j] = formatArticles[j], formatArticles[i]
	}
	return formatArticles, err
}

// GetArticle 获取一篇文章
func (article *Article) GetArticle() error {
	return config.GetDatabase().First(&article).Error
}

// CreateArticle 新增文章
func (article *Article) CreateArticle() error {
	return config.GetDatabase().Create(&article).Error
}

// UpdateArticle 更新文章
func (article *Article) UpdateArticle() error {
	return config.GetDatabase().Updates(&article).Error
}

// DeleteArticle 删除文章
func (article *Article) DeleteArticle() error {
	return config.GetDatabase().Delete(&article).Error
}

// GetArticlesByClass 按分类获取文章
func GetArticlesByClass(id int) ([]FormatArticle, error) {
	var formatArticles []FormatArticle
	var classes []Class
	class := Class{}
	class.ID = id
	subclasses, err := class.GetSubclasses()
	if err != nil {
		classes = append(classes, class)
	} else {
		classes = append(subclasses, class)
	}
	for _, c := range classes {
		var articles []Article
		err := config.GetDatabase().Where("class_id = ?", c.ID).Find(&articles).Error
		if err != nil {
			return formatArticles, err
		}
		for _, article := range articles {
			class := Class{}
			class.ID = article.ClassID
			err := class.GetClass()
			if err != nil {
				continue
			}
			formatArticle := FormatArticle{
				ID:          article.ID,
				Title:       article.Title,
				Class:       class.Name,
				CreatedDate: article.CreatedDate,
				UpdatedDate: article.UpdatedDate,
			}
			formatArticles = append(formatArticles, formatArticle)
		}
	}
	sort.Slice(formatArticles, func(i, j int) bool {
		return formatArticles[i].CreatedDate.After(formatArticles[j].CreatedDate)
	})
	return formatArticles, nil
}

// GetArticleNum 获取文章数
func GetArticleNum() int {
	var articles []Article
	err := config.GetDatabase().Find(&articles).Error
	if err != nil {
		return 0
	}
	return len(articles)
}
