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
	err := db.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	server := gin.Default()

	// JOB Routes
	server.POST("/api/jobs", handlers.CreateJOB)
	server.GET("/api/jobs", handlers.GetJobs)
	server.GET("/api/jobs/:id", handlers.GetJobByID)
	server.PUT("/api/jobs/:id", handlers.UpdateJob)
	server.DELETE("/api/jobs/:id", handlers.DeleteJob)
	// API Health
	server.GET("/healthz", handlers.Health)

	server.Run(":" + config.Port)

}
