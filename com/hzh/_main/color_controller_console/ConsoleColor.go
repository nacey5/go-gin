package color_controller_console

import "github.com/gin-gonic/gin"

//控制日志输出颜色

func ColorController() {
	// 禁止日志的颜色
	gin.DisableConsoleColor()

	// 用默认中间件创建一个 gin 路由:
	// 日志和恢复（无崩溃）中间件
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}

func ColorController2() {
	// 强制日志颜色化
	gin.ForceConsoleColor()

	// 用默认中间件创建一个 gin 路由:
	// 日志和恢复（无崩溃）中间件
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
