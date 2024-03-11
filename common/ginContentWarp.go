package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Render(c *gin.Context, tmp string, data gin.H) {
	c.HTML(http.StatusOK, tmp, data)
}

func Ok(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, msg)
}

func ServerErr(c *gin.Context, msg string) {
	c.JSON(http.StatusBadGateway, gin.H{
		"msg": msg,
	})
}

func BadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": msg,
	})
}
