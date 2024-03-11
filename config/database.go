package config

import (
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Charset  string
	//Local    string
}

var Database *gorm.DB

func GetDatabase() *gorm.DB {
	return Database
}
