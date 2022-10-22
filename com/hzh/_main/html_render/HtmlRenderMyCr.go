package html_render

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

// 自定义html模板渲染
func HtmlRenderMy() {
	router := gin.Default()
	//自定义分隔符
	router.Delims("{[{", "}]}")
	//回调方法
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("./templates/raw.tmpl")
	//请求路径
	router.GET("/raw", func(c *gin.Context) {
		//返回前端的数据类型为map
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	router.Run(":8080")
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
