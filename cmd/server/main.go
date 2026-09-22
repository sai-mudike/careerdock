package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/careerdock.git/internals/config"
	"github.com/sai-mudike/careerdock.git/internals/db"
	"github.com/sai-mudike/careerdock.git/internals/handlers"
	"github.com/sai-mudike/careerdock.git/internals/middleware"
)

func main() {

	config := config.MustLoad()
	err := db.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	server := gin.Default()
	server.MaxMultipartMemory = 5 << 20

	authenticator := server.Group("/")
	authenticator.Use(middleware.Authenticate)

	// JOB Routes
	authenticator.POST("/api/jobs", handlers.CreateJOB)
	authenticator.GET("/api/jobs", handlers.GetJobs)
	authenticator.GET("/api/jobs/:id", handlers.GetJobByID)
	authenticator.PUT("/api/jobs/:id", handlers.UpdateJob)
	authenticator.DELETE("/api/jobs/:id", handlers.DeleteJob)

	// Resume Routes
	authenticator.POST("/api/resumes", handlers.UploadResume)
	authenticator.GET("/api/resumes", handlers.GetResumes)
	authenticator.GET("/api/resumes/:id", handlers.GetResumeByID)
	authenticator.DELETE("/api/resumes/:id", handlers.DeleteResume)

	// User Routes
	server.POST("/api/register", handlers.RegisterUser)
	server.POST("/api/login", handlers.UserLogin)
	// API Health
	server.GET("/healthz", handlers.Health)

	server.Run(":" + config.Port)

}
