package services

import (
	"context"
	"net/http"
	"strings"

	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/repositories"
)

func UploadResume(ctx context.Context, userId string, resume models.ResumeRequest) (models.ResumeResponse, error) {

	validateResume(&resume)

	resumeFromClient := models.NewResume(userId, resume.Name, resume.FileName, resume.FilePath)

	resumeFromDB, err := repositories.CreateResume(ctx, *resumeFromClient)

	if err != nil {
		return models.ResumeResponse{}, err
	}

	resumeResponse := models.NewResumeResponse(resumeFromDB.Id, resumeFromDB.Name, resumeFromDB.FileName, resumeFromDB.FilePath, resumeFromDB.CreatedAT)

	return *resumeResponse, nil

}

func GetResumes(ctx context.Context, userID string) ([]models.ResumeResponse, error) {

	resumesFromDB, err := repositories.GetResumes(ctx, userID)

	if err != nil {
		return nil, err
	}
	return resumesFromDB, nil

}

func GetResumesByID(ctx context.Context, resumeID, userID string) (models.ResumeResponse, error) {

	resumeFromDB, err := repositories.GetResumeByID(ctx, resumeID, userID)

	if err != nil {
		return models.ResumeResponse{}, err
	}

	return resumeFromDB, nil
}

func DeleteResume(ctx context.Context, resumeID, Userid string) error {

	_, err := repositories.GetResumeByID(ctx, resumeID, Userid)

	if err != nil {
		return err
	}
	err = repositories.DeleteResume(ctx, resumeID, Userid)

	if err != nil {
		return err
	}

	return nil
}

func validateResume(resume *models.ResumeRequest) error {
	resume.Name = strings.TrimSpace(resume.Name)
	if resume.Name == "" {
		return customErr.New(customErr.CodeInvalidResumeName, "resume name is required", http.StatusBadRequest, nil)
	}

	if len(resume.Name) > 100 {
		return customErr.New(customErr.CodeInvalidResumeName, "resume name must not exceed 100 characters", http.StatusBadRequest, nil)
	}

	return nil
}
