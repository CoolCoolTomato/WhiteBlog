package main

import (
	"WhiteBlog/common"
	_ "WhiteBlog/init"
	"WhiteBlog/routes"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	router := gin.Default()
	routes.SetRoutes(router)
	router.Static("/static", "./static")
	router.Static("/uploads", "./uploads")
	router.SetFuncMap(template.FuncMap{
		"formatDateTime": common.FormatDateTime,
		"formatDate":     common.FormatDate,
	})
	router.LoadHTMLGlob("templates/**/*")
	router.Run(":7891")
}
