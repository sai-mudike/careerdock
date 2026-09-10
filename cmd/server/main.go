package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/careerdock.git/internals/config"
	"github.com/sai-mudike/careerdock.git/internals/db"
	"github.com/sai-mudike/careerdock.git/internals/handlers"
)

func main() {

	config := config.MustLoad()
	_, err := db.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	server := gin.Default()
	server.GET("/healthz", handlers.Health)
	server.Run(":" + config.Port)

}
