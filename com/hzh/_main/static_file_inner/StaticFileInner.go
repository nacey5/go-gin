package static_file_inner

//
//import (
//	"github.com/gin-gonic/gin"
//	"html/template"
//	"io/ioutil"
//	"net/http"
//	"strings"
//)
//
//// 静态资源嵌入
//func StaticFileInner() {
//	r := gin.New()
//
//	t, err := loadTemplate()
//	if err != nil {
//		panic(err)
//	}
//	r.SetHTMLTemplate(t)
//
//	r.GET("/", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "/html/index.tmpl", nil)
//	})
//	r.Run(":8080")
//}
//
//// loadTemplate 加载由 go-assets-builder 嵌入的模板
//func loadTemplate() (*template.Template, error) {
//	t := template.New("")
//	for name, file := range Assets.Files {
//		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
//			continue
//		}
//		h, err := ioutil.ReadAll(file)
//		if err != nil {
//			return nil, err
//		}
//		t, err = t.New(name).Parse(string(h))
//		if err != nil {
//			return nil, err
//		}
//	}
//	return t, nil
//}
