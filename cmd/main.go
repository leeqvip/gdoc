package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leeqvip/gdoc"
)

func main() {
	router := gin.Default()
	gdoc.UseInGin(router)
	router.Run(":9999")
}
