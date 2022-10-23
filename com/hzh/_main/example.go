package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/com/hzh/_main/redirect"
)

// 快速入门
func main() {
	redirect.RedirectGo()
}

func firstIn() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() //监听并在0.0.0.0上启动服务
}
