package set_get_cookie

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 设置和获取 Cookie
func SetAndGetCookie() {
	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}
