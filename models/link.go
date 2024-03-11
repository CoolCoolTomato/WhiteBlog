package models

import "WhiteBlog/config"

type Link struct {
	BaseModel
	Title       string
	Email       string
	Gravatar    string
	Url         string
	Description string
}

// GetLinks 获取所有友链
func GetLinks() ([]Link, error) {
	var links []Link
	err := config.GetDatabase().Find(&links).Error
	if err != nil {
		return links, err
	}
	return links, nil
}

// GetLink 获取某个友链
func (link *Link) GetLink() error {
	return config.GetDatabase().First(link).Error
}

// CreateLink 新增友链
func (link *Link) CreateLink() error {
	return config.GetDatabase().Create(link).Error
}

// UpdateLink 更新某个友链
func (link *Link) UpdateLink() error {
	return config.GetDatabase().Updates(&link).Error
}

// DeleteLink 删除某个友链
func (link *Link) DeleteLink() error {
	return config.GetDatabase().Delete(link).Error
}
