package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Inicialize() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on
}