package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/services"
)

func UploadResume(context *gin.Context) {

	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")
	const MaxResumeSize int64 = 5 << 20
	context.Request.Body = http.MaxBytesReader(
		context.Writer,
		context.Request.Body,
		MaxResumeSize,
	)

	file, err := context.FormFile("resume")

	if err != nil {

		HandleErrorWithGin(context, customErr.New(customErr.CodeFileRequired, "resume file is required", http.StatusBadRequest, err))
		return
	}

	if file.Size > MaxResumeSize {

		HandleErrorWithGin(context, customErr.New(customErr.CodeFileTooLarge, "resume must not exceed 5 MB", http.StatusRequestEntityTooLarge, err))
		return
	}

	dst := filepath.Join("./internals/assets/resumes", filepath.Base(file.Filename))

	isValidPDF, err := isPDF(file)

	if err != nil {
		HandleErrorWithGin(context, fmt.Errorf("Is pdf: %w", err))
		return
	}

	if !isValidPDF {
		HandleErrorWithGin(context, customErr.New(customErr.CodeInvalidFileType, "resume must be a PDF file", http.StatusBadRequest, err))
		return
	}
	err = context.SaveUploadedFile(file, dst)

	if err != nil {
		HandleErrorWithGin(context, fmt.Errorf("save pdf: %w", err))
		return
	}

	UserfileName := context.PostForm("filename")

	if UserfileName == "" {
		HandleErrorWithGin(context, customErr.New(customErr.CodeInvalidResumeName, "resume name is required", http.StatusBadRequest, nil))

		return
	}

	resumeFromClient := models.NewResumeRequest(UserfileName, file.Filename, dst)

	resumeFromDB, err := services.UploadResume(ctx, userIdFromContext, *resumeFromClient)

	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}

	context.JSON(http.StatusAccepted, resumeFromDB)
}

func GetResumes(context *gin.Context) {
	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")

	resumesFromDB, err := services.GetResumes(ctx, userIdFromContext)

	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}

	context.JSON(http.StatusOK, resumesFromDB)

}

func GetResumeByID(context *gin.Context) {
	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")
	resumeIDFromQuery := context.Param("id")

	if err := uuid.Validate(resumeIDFromQuery); err != nil {
		HandleErrorWithGin(context, customErr.New(
			customErr.CodeInvalidUUID,
			"invalid job id",
			http.StatusBadRequest,
			err,
		))

		return
	}
	resumeFromDB, err := services.GetResumesByID(ctx, resumeIDFromQuery, userIdFromContext)

	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}
	context.JSON(http.StatusOK, resumeFromDB)

}

func DeleteResume(context *gin.Context) {
	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")
	resumeIDFromQuery := context.Param("id")

	if err := uuid.Validate(resumeIDFromQuery); err != nil {
		HandleErrorWithGin(context, customErr.New(
			customErr.CodeInvalidUUID,
			"invalid job id",
			http.StatusBadRequest,
			err,
		))

		return
	}

	err := services.DeleteResume(ctx, resumeIDFromQuery, userIdFromContext)
	if err != nil {
		HandleErrorWithGin(context, err)
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "resume delete successfully"})
}

func isPDF(file *multipart.FileHeader) (bool, error) {

	if strings.ToLower(filepath.Ext(file.Filename)) != ".pdf" {
		return false, fmt.Errorf("invalid file extention")
	}

	f, err := file.Open()

	if err != nil {
		return false, err
	}

	defer f.Close()

	buffer := make([]byte, 512)

	n, err := f.Read(buffer)

	contentType := http.DetectContentType(buffer[:n])

	return contentType == "application/pdf", err

}
