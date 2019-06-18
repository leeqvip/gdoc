package gdoc

import (
	"github.com/gin-gonic/gin"
)

func UseInGin(g *gin.Engine)  {
	g.GET("/docs/*name", func(context *gin.Context) {
		body := Handler(context.Param("name"))
		context.Header("Content-Type", "text/html; charset=utf-8")
		context.String(200, string(body))
	})
}