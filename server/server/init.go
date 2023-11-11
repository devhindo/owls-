package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {

	router := gin.Default()

	router.LoadHTMLGlob("views/html/*")
	router.Static("/views/css", "./css")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	router.GET("/chat", func(c *gin.Context) {
        c.String(200, "Welcome to the chat!")
    })

	router.Run(":8080")
}