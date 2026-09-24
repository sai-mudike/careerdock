package services

import (
	"context"
	"net/http"
	"strings"

	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/repositories"
)

func CreateApplication(ctx context.Context, userID string, application models.ApplicationRequest) (models.ApplicationResponse, error) {

	_, err := repositories.GetJobByID(ctx, application.JobID, userID)

	if err != nil {
		return models.ApplicationResponse{}, err
	}

	_, err = repositories.GetResumeByID(ctx, application.ResumeID, userID)

	if err != nil {
		return models.ApplicationResponse{}, err
	}

	applicationFromClient := models.NewApplication(userID, application.JobID, application.ResumeID, application.Notes)

	applicationFromDB, err := repositories.CreateApplication(ctx, *applicationFromClient)

	if err != nil {
		return models.ApplicationResponse{}, err
	}

	applicationResponse := models.NewApplicationResponse(applicationFromDB.Id, applicationFromDB.JobID, applicationFromDB.ResumeID, applicationFromDB.Status, applicationFromDB.Notes, applicationFromDB.AppliedAT, applicationFromDB.CreatedAT, applicationFromDB.UpdatedAT)

	return *applicationResponse, nil

}

func GetAllApplications(ctx context.Context, userId string) ([]models.ApplicationResponse, error) {

	applicationsFromDB, err := repositories.GetAllApplications(ctx, userId)

	if err != nil {
		return nil, err
	}

	return applicationsFromDB, nil
}

func GetApplicationByID(ctx context.Context, applicationID, userID string) (models.ApplicationResponse, error) {

	applicationFromDB, err := repositories.GetApplicationByID(ctx, applicationID, userID)

	if err != nil {
		return models.ApplicationResponse{}, err
	}

	return applicationFromDB, nil

}

func UpdateApplication(ctx context.Context, applicationId, userId string, application models.UpdateApplicationRequest) (models.ApplicationResponse, error) {
	_, err := repositories.GetApplicationByID(ctx, applicationId, userId)

	if err != nil {
		return models.ApplicationResponse{}, err
	}

	if application.ResumeID != nil {

		_, err = repositories.GetResumeByID(ctx, *application.ResumeID, userId)

		if err != nil {
			return models.ApplicationResponse{}, err
		}

	}

	*application.Status = strings.ToLower(strings.TrimSpace(*application.Status))

	if err := validateApplicationStatus(*application.Status); err != nil {
		return models.ApplicationResponse{}, err
	}

	if application.Notes != nil {
		*application.Notes = strings.TrimSpace(*application.Notes)
	}

	applicationFromClient := models.Application{
		Id:        applicationId,
		UserID:    userId,
		ResumeID:  *application.ResumeID,
		Status:    *application.Status,
		Notes:     *application.Notes,
		AppliedAT: application.AppliedAT,
	}

	applicationFromDB, err := repositories.UpdateApplications(ctx, applicationFromClient)
	if err != nil {
		return models.ApplicationResponse{}, err
	}

	applicationResponse := models.NewApplicationResponse(applicationFromDB.Id, applicationFromDB.JobID, applicationFromDB.ResumeID, applicationFromDB.Status, applicationFromDB.Notes, applicationFromDB.AppliedAT, applicationFromDB.CreatedAT, applicationFromDB.UpdatedAT)

	return *applicationResponse, nil

}

func DeleteApplication(ctx context.Context, applicationID, userID string) error {

	_, err := repositories.GetApplicationByID(ctx, applicationID, userID)

	if err != nil {
		return err
	}

	err = repositories.DeleteApplication(ctx, applicationID, userID)

	if err != nil {
		return err
	}

	return nil

}

func validateApplicationStatus(status string) error {

	switch status {
	case "applied",
		"interview",
		"offer",
		"rejected",
		"withdrawn",
		"accepted":
		return nil

	default:
		return customErr.New(customErr.CodeInvalidApplicationStatus, "invalid application status", http.StatusBadRequest, nil)
	}
}
