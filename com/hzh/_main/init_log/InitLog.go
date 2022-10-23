package init_log

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 默认的日志格式
// [GIN-debug] POST   /foo                      --> main.main.func1 (3 handlers)
// [GIN-debug] GET    /bar                      --> main.main.func2 (3 handlers)
// [GIN-debug] GET    /status                   --> main.main.func3 (3 handlers)
// 自定义日志格式
func InitLog() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	r.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run()
}
