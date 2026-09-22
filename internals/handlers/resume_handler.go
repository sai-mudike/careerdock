package handlers

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
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
		HandleError(context, err)
		return
	}

	if file.Size > MaxResumeSize {
		HandleError(context, customErr.ErrMaxResumeSize)
		return
	}

	dst := filepath.Join("./internals/assets/resumes", filepath.Base(file.Filename))

	isValidPDF, err := isPDF(file)

	if !isValidPDF {
		HandleError(context, err)
		return
	}
	err = context.SaveUploadedFile(file, dst)

	if err != nil {
		HandleError(context, err)
		return
	}

	UserfileName := context.PostForm("filename")

	if UserfileName == "" {
		HandleError(context, customErr.ErrResumeName)
		return
	}

	resumeFromClient := models.NewResumeRequest(UserfileName, file.Filename, dst)

	resumeFromDB, err := services.UploadResume(ctx, userIdFromContext, *resumeFromClient)

	if err != nil {
		HandleError(context, err)
		return
	}

	context.JSON(http.StatusAccepted, resumeFromDB)
}

func GetResumes(context *gin.Context) {
	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")

	resumesFromDB, err := services.GetResumes(ctx, userIdFromContext)

	if err != nil {
		HandleError(context, err)
		return
	}

	context.JSON(http.StatusOK, resumesFromDB)

}

func GetResumeByID(context *gin.Context) {
	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")
	resumeIDFromQuery := context.Param("id")
	resumeFromDB, err := services.GetResumesByID(ctx, resumeIDFromQuery, userIdFromContext)

	if err != nil {
		HandleError(context, err)
		return
	}
	context.JSON(http.StatusOK, resumeFromDB)

}

func DeleteResume(context *gin.Context) {
	ctx := context.Request.Context()
	userIdFromContext := context.GetString("userID")
	resumeIDFromQuery := context.Param("id")

	err := services.DeleteResume(ctx, resumeIDFromQuery, userIdFromContext)
	if err != nil {
		HandleError(context, err)
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "resume delete successfully"})
}

func isPDF(file *multipart.FileHeader) (bool, error) {

	if strings.ToLower(filepath.Ext(file.Filename)) != ".pdf" {
		return false, customErr.ErrInvalidResumeData
	}

	f, err := file.Open()

	if err != nil {
		return false, err
	}

	defer f.Close()

	buffer := make([]byte, 512)

	n, err := f.Read(buffer)

	contentType := http.DetectContentType(buffer[:n])

	return contentType == "application/pdf", nil

}
