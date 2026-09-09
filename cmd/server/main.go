package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/careerdock.git/internals/config"
)

func main() {

	config := config.MustLoad()

	server := gin.Default()
	server.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	server.Run(":" + config.Port)

}
