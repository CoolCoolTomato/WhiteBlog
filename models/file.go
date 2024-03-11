package models

import "WhiteBlog/config"

type File struct {
	ID       int `gorm:"primary_key"`
	FileName string
	FilePath string
}

// GetFiles 获取文件列表
func GetFiles() ([]File, error) {
	var files []File
	err := config.GetDatabase().Find(&files).Error
	return files, err
}

// GetFile 获取文件
func (file *File) GetFile() error {
	return config.GetDatabase().First(&file).Error
}

// AddFile 添加文件
func (file *File) AddFile() error {
	return config.GetDatabase().Create(&file).Error
}

// UpdateFile 更新文件
func (file *File) UpdateFile() error {
	return config.GetDatabase().Updates(&file).Error
}

// DeleteFile 删除文件
func (file *File) DeleteFile() error {
	return config.GetDatabase().Delete(&file).Error
}
