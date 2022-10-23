package redirect

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RedirectGo() {

	r := gin.Default()
	//get重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})
	//get重定向
	r.POST("/test", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})

	//路由器重定向
	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/test3"
		r.HandleContext(c)
	})
	r.GET("/test3", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
	r.Run(":8080")
}
