package models

import (
	"WhiteBlog/config"
)

type Poet struct {
	BaseModel
	Content string
	Source  string
}

func GetPoets() ([]Poet, error) {
	var poets []Poet
	err := config.GetDatabase().Find(&poets).Error
	return poets, err
}

// GetPoet 获取一条句子
func (poet *Poet) GetPoet() error {
	return config.GetDatabase().First(&poet).Error
}

// CreatePoet 新增句子
func (poet *Poet) CreatePoet() error {
	return config.GetDatabase().Create(&poet).Error
}

// UpdatePoet 修改句子
func (poet *Poet) UpdatePoet() error {
	return config.GetDatabase().Updates(&poet).Error
}

// DeletePoet 删除句子
func (poet *Poet) DeletePoet() error {
	return config.GetDatabase().Delete(&poet).Error
}
