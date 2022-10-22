package ascii_json

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Example() {
	r := gin.Default()
	r.GET("/someJSON", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "<br>",
		}

		//输出: 输出状体以及ASCII对应的码:{"lang":"Go\u8bed\u8a00","tag":"\u003cbr\u003e"}
		context.AsciiJSON(http.StatusOK, data)
	})

	r.Run()

}
