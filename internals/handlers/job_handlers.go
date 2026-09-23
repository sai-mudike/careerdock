package handlers

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/services"
	"github.com/sai-mudike/careerdock.git/internals/validation"
)

func CreateJOB(context *gin.Context) {

	var jobFromClient models.JobRequest
	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")

	err := context.ShouldBindJSON(&jobFromClient)

	if err != nil {
		HandleErrorWithGin(context, customErr.New(
			"INVALID_REQUEST",
			validation.ValidationMessage(err),
			http.StatusBadRequest,
			err,
		))
		return
	}

	jobFromDB, err := services.CreateJOB(ctx, userIdFromContext, jobFromClient)

	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}

	context.JSON(http.StatusCreated, jobFromDB)

}

func GetJobs(context *gin.Context) {
	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")

	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil {

		HandleErrorWithGin(context, customErr.New(
			customErr.CodeInvalidPagination,
			"page must be a number",
			http.StatusBadRequest,
			err,
		))
		return
	}

	limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
	if err != nil {
		HandleErrorWithGin(context, customErr.New(
			customErr.CodeInvalidPagination,
			"limit must be a number",
			http.StatusBadRequest,
			err,
		))
		return
	}

	query := models.JobQuery{
		Page:    page,
		Limit:   limit,
		SortBy:  context.DefaultQuery("sortby", "created_at"),
		OrderBy: context.DefaultQuery("orderby", "desc"),
	}

	jobs, err := services.GetAllJobs(ctx, userIdFromContext, query)

	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}
	context.JSON(http.StatusOK, jobs)

}

func GetJobByID(context *gin.Context) {
	ctx := context.Request.Context()
	jobID := context.Param("id")
	userIdFromContext := context.GetString("userID")

	if err := uuid.Validate(jobID); err != nil {
		HandleErrorWithGin(context, customErr.New(
			customErr.CodeInvalidUUID,
			"invalid job id",
			http.StatusBadRequest,
			err,
		))

		return
	}

	job, err := services.GetJobByID(ctx, jobID, userIdFromContext)
	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}

	context.JSON(http.StatusOK, job)
}

func UpdateJob(context *gin.Context) {
	jobID := context.Param("id")
	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")
	if err := uuid.Validate(jobID); err != nil {
		HandleErrorWithGin(context, customErr.New(
			customErr.CodeInvalidUUID,
			"invalid job id",
			http.StatusBadRequest,
			err,
		))

		return
	}

	var jobFromClient models.JobRequest
	err := context.ShouldBindJSON(&jobFromClient)
	if err != nil {
		HandleErrorWithGin(context, customErr.New(
			"INVALID_REQUEST",
			err.Error(),
			http.StatusBadRequest,
			err,
		))
		return
	}

	updatedJOB, err := services.UpdateJob(ctx, jobID, userIdFromContext, jobFromClient)
	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}

	context.JSON(http.StatusOK, updatedJOB)

}

func DeleteJob(context *gin.Context) {
	ctx := context.Request.Context()
	jobID := context.Param("id")
	userIdFromContext := context.GetString("userID")
	if err := uuid.Validate(jobID); err != nil {
		HandleErrorWithGin(context, customErr.New(
			customErr.CodeInvalidUUID,
			"invalid job id",
			http.StatusBadRequest,
			err,
		))

		return
	}

	err := services.DeleteJob(ctx, jobID, userIdFromContext)

	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "job deleted succefully"})
}
