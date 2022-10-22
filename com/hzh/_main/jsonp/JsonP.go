package jsonp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonP() {
	r := gin.Default()
	r.GET("/JSONP", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		//JSONP?callback=x  {"foo":"bar"}
		context.JSONP(http.StatusOK, data)
	})
	r.Run(":8080")
}
