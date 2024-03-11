package admin

import (
	"WhiteBlog/common"
	"WhiteBlog/config"
	"WhiteBlog/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {
	theConfig := config.GetConfig()
	if c.Request.Method == "POST" {
		// 从请求中获取用户名和密码
		username := c.PostForm("username")
		password := c.PostForm("password")
		scode := c.PostForm("scode")
		// 验证用户名密码
		if !verify(username, password, scode) {
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}
		session := sessions.Default(c)
		session.Options(sessions.Options{
			MaxAge: 43200,
		})
		session.Set("user", "admin")
		_ = session.Save()
		c.Redirect(http.StatusSeeOther, "/admin")
		return
	}
	common.Render(c, "admin/login.html", gin.H{
		"config": theConfig,
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	_ = session.Save()
	c.Redirect(http.StatusSeeOther, "/login")
}

func verify(username string, password string, scode string) bool {
	code, _ := strconv.Atoi(scode)
	if username == config.GetConfig().AuthConfig.Username && password == config.GetConfig().AuthConfig.Password {
		if config.GetConfig().AuthConfig.Enable2FA {
			if !middleware.VerifyCode(common.StrTo16Upper(config.GetConfig().AuthConfig.Username), int32(code)) {
				return false
			}
		}
		return true
	}
	return false
}
