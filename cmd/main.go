package main

import (
	"github.com/gin-gonic/gin"
	"github.com/techoner/gindoc"
)

func main()  {
	router := gin.Default()
	gindoc.UseInGin(router)
	router.Run(":9999")
}