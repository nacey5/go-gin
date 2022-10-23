package ginNotDefault

import "github.com/gin-gonic/gin"

func init() {
	// Default 使用 Logger 和 Recovery 中间件
	r := gin.Default()
	println(r)
	//如果不使用默认的中间件的话
	r = gin.New()
	println(r)
}
