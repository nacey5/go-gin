package query_postform

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// POST /post?id=1234&page=1 HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
//
// name=manu&message=this_is_great
func QueryForm() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		//路径查询
		id := c.Query("id")
		//如果没page参数，则默认为1，有的话按有的来进行计算
		page := c.DefaultQuery("page", "0")
		//表单查询
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}
