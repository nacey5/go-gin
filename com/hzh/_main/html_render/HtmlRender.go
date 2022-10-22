package html_render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//html渲染

func HtmlRender() {
	router := gin.Default()
	//识别templates包下的所有文件
	//bug记录:当读取的路径包含子包的情况下，会出现初始化异常，需要读取所有文件-
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(context *gin.Context) {
		//将某个文件识别为html格式
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}

func HtmlRender2() {
	router := gin.Default()
	//识别包路径
	router.LoadHTMLGlob("templates/**/*")
	//访问路径对应的资源
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	//访问路径对应的资源
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.Run(":8080")
}
