package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/services"
)

func CreateJOB(context *gin.Context) {

	var jobFromClient models.JobRequest
	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")

	err := context.ShouldBindJSON(&jobFromClient)

	if err != nil {
		HandleError(context, customErr.ErrInvalidJobData)
		return
	}

	jobFromDB, err := services.CreateJOB(ctx, userIdFromContext, jobFromClient)

	if err != nil {
		HandleError(context, err)

		return
	}

	context.JSON(http.StatusCreated, jobFromDB)

}

func GetJobs(context *gin.Context) {
	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")

	jobs, err := services.GetAllJobs(ctx, userIdFromContext)

	if err != nil {
		HandleError(context, err)
		return
	}
	context.JSON(http.StatusOK, jobs)

}

func GetJobByID(context *gin.Context) {
	ctx := context.Request.Context()
	jobID := context.Param("id")
	userIdFromContext := context.GetString("userID")

	job, err := services.GetJobByID(ctx, jobID, userIdFromContext)
	if err != nil {
		HandleError(context, err)
		return
	}

	context.JSON(http.StatusOK, job)
}

func UpdateJob(context *gin.Context) {
	jobID := context.Param("id")
	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")

	var jobFromClient models.JobRequest
	err := context.ShouldBindJSON(&jobFromClient)
	if err != nil {
		HandleError(context, customErr.ErrInvalidJobData)
		return
	}

	updatedJOB, err := services.UpdateJob(ctx, jobID, userIdFromContext, jobFromClient)
	if err != nil {
		HandleError(context, err)

		return
	}

	context.JSON(http.StatusAccepted, updatedJOB)

}

func DeleteJob(context *gin.Context) {
	ctx := context.Request.Context()
	jobID := context.Param("id")
	userIdFromContext := context.GetString("userID")

	err := services.DeleteJob(ctx, jobID, userIdFromContext)

	if err != nil {
		HandleError(context, err)
		return
	}

	context.JSON(http.StatusNotFound, gin.H{"message": "job deleted succefully"})
}
