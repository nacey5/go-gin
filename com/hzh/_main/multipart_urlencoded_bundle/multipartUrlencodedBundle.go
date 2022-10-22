package multipart_urlencoded_bundle

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 设计表单
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func MUB() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var form LoginForm
		// 在这种情况下，将自动选择合适的绑定
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	router.Run(":8080")
}

func MUB2() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// 使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		var form LoginForm
		if c.ShouldBindWith(&form, binding.Form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	router.Run(":8080")
}
