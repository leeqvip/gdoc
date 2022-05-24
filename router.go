package gdoc

import (
	"github.com/gin-gonic/gin"
)

func UseInGin(r gin.IRouter)  {
	r.GET("/docs/*name", func(context *gin.Context) {
		body := Handle(context.Param("name"))
		context.Header("Content-Type", "text/html; charset=utf-8")
		context.String(200, string(body))
	})
}
