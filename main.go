package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Merhaba Dünya! Bugün nasılsın")
	})
	router.Run(":8080")
}
