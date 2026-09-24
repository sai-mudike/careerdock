package handlers

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/services"
	"github.com/sai-mudike/careerdock.git/internals/validation"
)

func CreateApplication(context *gin.Context) {

	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")

	if err := uuid.Validate(userIdFromContext); err != nil {
		HandleErrorWithGin(context, customErr.New(customErr.CodeInvalidUUID, "invalid user id", http.StatusBadRequest, err))
		return

	}

	var applicationFromClient models.ApplicationRequest

	err := context.ShouldBindJSON(&applicationFromClient)

	if err != nil {
		HandleErrorWithGin(context, customErr.New(customErr.CodeInvalidApplicationData, validation.ValidationMessage(err), http.StatusBadRequest, err))
		return
	}

	if err := uuid.Validate(applicationFromClient.JobID); err != nil {
		HandleErrorWithGin(context, customErr.New(customErr.CodeInvalidUUID, "invalid job id", http.StatusBadRequest, err))
		return

	}
	if err := uuid.Validate(applicationFromClient.ResumeID); err != nil {
		HandleErrorWithGin(context, customErr.New(customErr.CodeInvalidUUID, "invalid resume id", http.StatusBadRequest, err))
		return

	}

	applicationFromDB, err := services.CreateApplication(ctx, userIdFromContext, applicationFromClient)

	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}

	context.JSON(http.StatusCreated, applicationFromDB)

}

func GetAllApplications(context *gin.Context) {
	userIDFromContext := context.GetString("userID")
	ctx := context.Request.Context()

	applicationsFromDB, err := services.GetAllApplications(ctx, userIDFromContext)

	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}
	context.JSON(http.StatusOK, applicationsFromDB)

}

func GetApplicationByID(context *gin.Context) {
	userIDFromContext := context.GetString("userID")
	applicationIdFromQuery := context.Param("id")
	ctx := context.Request.Context()

	if err := uuid.Validate(applicationIdFromQuery); err != nil {
		HandleErrorWithGin(context, customErr.New(customErr.CodeInvalidUUID, "invalid job id", http.StatusBadRequest, err))
		return
	}

	applicationFromDB, err := services.GetApplicationByID(ctx, applicationIdFromQuery, userIDFromContext)

	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}
	context.JSON(http.StatusOK, applicationFromDB)

}

func UpdateApplication(context *gin.Context) {
	userIDFromContext := context.GetString("userID")
	applicationIdFromQuery := context.Param("id")
	ctx := context.Request.Context()

	if err := uuid.Validate(applicationIdFromQuery); err != nil {
		HandleErrorWithGin(context, customErr.New(customErr.CodeInvalidUUID, "invalid application id", http.StatusBadRequest, err))
		return
	}

	var applicationFromClient models.UpdateApplicationRequest

	err := context.ShouldBindJSON(&applicationFromClient)

	if err != nil {
		HandleErrorWithGin(context, customErr.New(customErr.CodeInvalidApplicationData, validation.ValidationMessage(err), http.StatusBadRequest, err))
		return
	}
	if err := uuid.Validate(*applicationFromClient.ResumeID); err != nil {
		HandleErrorWithGin(context, customErr.New(customErr.CodeInvalidUUID, "invalid resume id", http.StatusBadRequest, err))
		return
	}
	applicationFromDB, err := services.UpdateApplication(ctx, applicationIdFromQuery, userIDFromContext, applicationFromClient)

	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}

	context.JSON(http.StatusOK, applicationFromDB)

}

func DeleteApplication(context *gin.Context) {
	userIDFromContext := context.GetString("userID")
	applicationIdFromQuery := context.Param("id")
	ctx := context.Request.Context()

	if err := uuid.Validate(applicationIdFromQuery); err != nil {
		HandleErrorWithGin(context, customErr.New(customErr.CodeInvalidUUID, "invalid job id", http.StatusBadRequest, err))
		return
	}

	err := services.DeleteApplication(ctx, applicationIdFromQuery, userIDFromContext)

	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "application deleted successfully"})
}
