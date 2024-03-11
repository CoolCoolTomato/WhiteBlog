package admin

import (
	"WhiteBlog/common"
	"WhiteBlog/config"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
)

func Setting(c *gin.Context) {
	theConfig := config.GetConfig()
	if c.Request.Method == "POST" {
		err := c.ShouldBindJSON(&config.TheConfig)
		if err != nil {
			common.BadRequest(c, "Data Error")
			return
		}
		configBytes, err := json.MarshalIndent(&config.TheConfig, "", "    ")
		if err != nil {
			common.ServerErr(c, "Server Error")
			return
		}
		jsonFile, err := os.Create("config.json")
		if err != nil {
			common.ServerErr(c, "Server Error")
			return
		}
		defer jsonFile.Close()
		jsonFile.Write(configBytes)
		common.Ok(c, "Success")
		return
	}
	key := common.StrTo16Upper(config.GetConfig().AuthConfig.Username)
	common.Render(c, "admin/setting.html", gin.H{
		"config": theConfig,
		"key":    key,
	})
}
