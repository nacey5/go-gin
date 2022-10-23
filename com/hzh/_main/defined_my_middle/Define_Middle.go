package defined_my_middle

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Defined_Middle() {
	r := gin.New()
	//使用Logger这个中间件
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// 打印："12345"
		log.Println(example)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		//输出时间戳
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}
