package init

import (
	"WhiteBlog/config"
	"WhiteBlog/models"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func init() {
	jsonConfig, err := os.Open("config.json")
	if err != nil {
		return
	}
	defer jsonConfig.Close()
	err = json.NewDecoder(jsonConfig).Decode(&config.TheConfig)
	//  数据库
	databaseConfig := config.GetConfig().DatabaseConfig
	switch config.GetConfig().DatabaseConfig.Driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
			databaseConfig.Username,
			databaseConfig.Password,
			databaseConfig.Host,
			databaseConfig.Port,
			databaseConfig.Database,
			databaseConfig.Charset)
		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}
		config.Database = conn
	}
	//  迁移表
	config.Database.AutoMigrate(
		&models.Class{},
		&models.Article{},
		&models.Image{},
		&models.Link{},
		&models.Poet{},
		&models.File{},
	)
}
