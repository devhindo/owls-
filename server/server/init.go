package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {

	router := gin.Default()

	router.LoadHTMLGlob("views/html/*")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	router.Run(":8080")
}