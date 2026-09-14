package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/repositories"
)

func CreateJOB(ctx context.Context, job models.Job) (models.Job, error) {
	uui, err := uuid.Parse("8e6c7811-9a7c-47b7-85b8-5c02b9e11c52")
	if err != nil {
		return models.Job{}, customErr.ErrInvalidRequest
	}

	job.UserID = uui

	jobFromDB, err := repositories.CreateJob(ctx, job)
	if err != nil {
		return models.Job{}, err
	}

	return jobFromDB, nil
}

func GetAllJobs(ctx context.Context) ([]models.Job, error) {

	jobsFromDB, err := repositories.GetAllJobs(ctx)

	if err != nil {
		return nil, err
	}

	return jobsFromDB, nil

}

func GetJobByID(ctx context.Context, jobID string) (models.Job, error) {
	parsedUUID, err := uuid.Parse(jobID)

	if err != nil {
		return models.Job{}, customErr.ErrInvalidRequest
	}

	jobsFromDB, err := repositories.GetJobByID(ctx, parsedUUID)

	if err != nil {
		return models.Job{}, err
	}

	return jobsFromDB, nil
}

func UpdateJob(ctx context.Context, jobID string, job *models.Job) (models.Job, error) {

	parsedUUID, err := uuid.Parse(jobID)
	if err != nil {
		return models.Job{}, customErr.ErrInvalidRequest
	}

	job.Id = parsedUUID

	_, err = repositories.GetJobByID(ctx, parsedUUID)
	if err != nil {
		return models.Job{}, err
	}

	updatedJob, err := repositories.UpdateJob(ctx, *job)
	if err != nil {
		return models.Job{}, err
	}

	return updatedJob, nil

}

func DeleteJob(ctx context.Context, jobID string) error {
	parsedUUID, err := uuid.Parse(jobID)

	if err != nil {
		return customErr.ErrInvalidRequest
	}

	_, err = repositories.GetJobByID(ctx, parsedUUID)
	if err != nil {
		return err
	}

	err = repositories.DeleteJob(ctx, parsedUUID)

	if err != nil {
		return err
	}
	return nil

}
