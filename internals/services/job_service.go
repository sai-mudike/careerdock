package services

import (
	"context"

	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/repositories"
)

func CreateJOB(ctx context.Context, job models.Job) (models.Job, error) {

	jobFromDB, err := repositories.CreateJob(ctx, job)
	if err != nil {
		return models.Job{}, err
	}

	return jobFromDB, nil
}

func GetAllJobs(ctx context.Context, userID string) ([]models.Job, error) {

	jobsFromDB, err := repositories.GetAllJobs(ctx, userID)

	if err != nil {
		return nil, err
	}

	return jobsFromDB, nil

}

func GetJobByID(ctx context.Context, jobID string, userID string) (models.Job, error) {

	jobsFromDB, err := repositories.GetJobByID(ctx, jobID, userID)

	if err != nil {
		return models.Job{}, err
	}

	return jobsFromDB, nil
}

func UpdateJob(ctx context.Context, jobID string, job models.Job) (models.Job, error) {

	jobFromDB, err := repositories.GetJobByID(ctx, jobID, job.UserID)
	if err != nil {
		return models.Job{}, err
	}

	if jobFromDB.UserID != job.UserID {
		return models.Job{}, customErr.ErrUnauthorized

	}

	updatedJob, err := repositories.UpdateJob(ctx, jobID, job)
	if err != nil {
		return models.Job{}, err
	}

	return updatedJob, nil

}

func DeleteJob(ctx context.Context, jobID, userID string) error {

	jobFromDB, err := repositories.GetJobByID(ctx, jobID, userID)
	if err != nil {
		return err
	}
	if jobFromDB.UserID != userID {
		return customErr.ErrUnauthorized

	}

	err = repositories.DeleteJob(ctx, jobID)

	if err != nil {
		return err
	}
	return nil

}
