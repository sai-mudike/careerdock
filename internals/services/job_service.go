package services

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/repositories"
)

func CreateJOB(ctx context.Context, userID string, job models.JobRequest) (models.JobResponse, error) {

	jobFromClient := models.NewJob(userID, job.CompanyName, job.Position, job.JobURL, job.Location, job.EmploymentType, job.SalaryMin, job.SalaryMax, job.Description)

	err := validateJob(jobFromClient)
	if err != nil {
		return models.JobResponse{}, err
	}

	jobFromDB, err := repositories.CreateJob(ctx, *jobFromClient)
	if err != nil {
		return models.JobResponse{}, err
	}

	jobResponceDTO := models.NewJobResponse(jobFromDB.Id, jobFromDB.CompanyName, jobFromDB.Position, jobFromDB.JobURL, jobFromDB.Location, jobFromDB.EmploymentType, jobFromDB.SalaryMin, jobFromDB.SalaryMax, jobFromDB.Description, jobFromDB.CreatedAT, jobFromDB.UpdatedAT)

	return *jobResponceDTO, nil
}

func GetAllJobs(ctx context.Context, userID string, query models.JobQuery) ([]models.JobResponse, error) {

	err := validateQuery(&query)

	if err != nil {
		return nil, err
	}

	jobsFromDB, err := repositories.GetAllJobs(ctx, userID, query)

	if err != nil {
		return nil, err
	}

	return jobsFromDB, nil

}

func GetJobByID(ctx context.Context, jobID string, userID string) (models.JobResponse, error) {

	jobsFromDB, err := repositories.GetJobByID(ctx, jobID, userID)

	if err != nil {
		return models.JobResponse{}, err
	}

	return jobsFromDB, nil
}

func UpdateJob(ctx context.Context, jobID, userId string, job models.JobRequest) (models.JobResponse, error) {

	_, err := repositories.GetJobByID(ctx, jobID, userId)
	if err != nil {
		return models.JobResponse{}, err
	}

	jobFromClient := models.NewJob(userId, job.CompanyName, job.Position, job.JobURL, job.Location, job.EmploymentType, job.SalaryMin, job.SalaryMax, job.Description)

	err = validateJob(jobFromClient)
	if err != nil {
		return models.JobResponse{}, err
	}

	jobFromDB, err := repositories.UpdateJob(ctx, jobID, *jobFromClient)
	if err != nil {
		return models.JobResponse{}, err
	}

	jobResponceDTO := models.NewJobResponse(jobFromDB.Id, jobFromDB.CompanyName, jobFromDB.Position, jobFromDB.JobURL, jobFromDB.Location, jobFromDB.EmploymentType, jobFromDB.SalaryMin, jobFromDB.SalaryMax, jobFromDB.Description, jobFromDB.CreatedAT, jobFromDB.UpdatedAT)

	return *jobResponceDTO, nil

}

func DeleteJob(ctx context.Context, jobID, userID string) error {

	_, err := repositories.GetJobByID(ctx, jobID, userID)
	if err != nil {
		return err
	}

	err = repositories.DeleteJob(ctx, jobID, userID)

	if err != nil {
		return err
	}
	return nil

}
func validateJob(job *models.Job) error {

	job.CompanyName = strings.ToLower(strings.TrimSpace(job.CompanyName))
	job.Position = strings.ToLower(strings.TrimSpace(job.Position))
	job.JobURL = strings.TrimSpace(job.JobURL)

	job.Location = strings.ToLower(strings.TrimSpace(job.Location))
	job.EmploymentType = strings.ToLower(strings.TrimSpace(job.EmploymentType))
	job.Description = strings.TrimSpace(job.Description)

	if job.CompanyName == "" {
		return customErr.New(customErr.CodeInvalidJobData, "company_name is required", http.StatusBadRequest, nil)
	}

	if len(job.CompanyName) > 250 {
		return customErr.New(
			customErr.CodeInvalidJobData,
			"company_name must not exceed 250 characters",
			http.StatusBadRequest,
			nil,
		)
	}

	if job.Position == "" {
		return customErr.New(
			customErr.CodeInvalidJobData,
			"position is required",
			http.StatusBadRequest,
			nil,
		)
	}

	if len(job.Position) > 250 {
		return customErr.New(
			customErr.CodeInvalidJobData,
			"position must not exceed 250 characters",
			http.StatusBadRequest,
			nil,
		)
	}

	if job.JobURL == "" {
		return customErr.New(
			customErr.CodeInvalidJobURL,
			"job_url is required",
			http.StatusBadRequest,
			nil,
		)
	}

	u, err := url.Parse(job.JobURL)

	if err != nil {
		return customErr.New(
			customErr.CodeInvalidJobURL,
			"job_url must be a valid URL",
			http.StatusBadRequest,
			nil,
		)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return customErr.New(
			customErr.CodeInvalidJobURL,
			"job_url must use http or https",
			http.StatusBadRequest,
			nil,
		)
	}

	if job.EmploymentType == "" {
		return customErr.New(
			customErr.CodeInvalidJobData,
			"invalid employment type",
			http.StatusBadRequest,
			nil,
		)
	}

	if job.SalaryMin != 0 && job.SalaryMin < 0 {
		return customErr.New(
			customErr.CodeInvalidJobData,
			"salary cannot be negative",
			http.StatusBadRequest,
			nil,
		)
	}

	if job.SalaryMax != 0 && job.SalaryMax < 0 {
		return customErr.New(
			customErr.CodeInvalidJobData,
			"salary cannot be negative",
			http.StatusBadRequest,
			nil,
		)
	}
	// Salary range
	if job.SalaryMin != 0 &&
		job.SalaryMax != 0 &&
		job.SalaryMin > job.SalaryMax {

		return customErr.New(
			customErr.CodeInvalidJobData,
			"salary_min cannot be greater than salary_max",
			http.StatusBadRequest,
			nil,
		)
	}
	// Notes
	if len(job.Description) > 500 {
		return customErr.New(
			customErr.CodeInvalidJobData,
			"description must not exceed 500 characters",
			http.StatusBadRequest,
			nil,
		)
	}

	return nil
}

func validateQuery(query *models.JobQuery) error {

	allowedSorts := map[string]bool{
		"created_at":   true,
		"updated_at":   true,
		"company_name": true,
		"salary_min":   true,
		"salary_max":   true,
	}

	if query.Page <= 0 {
		return customErr.New(
			customErr.CodeInvalidPagination,
			"page must be greater than 0",
			http.StatusBadRequest,
			nil,
		)
	}

	if query.Limit <= 0 || query.Limit >= 100 {
		return customErr.New(
			customErr.CodeInvalidPagination,
			"limit must be between 1 and 100",
			http.StatusBadRequest,
			nil,
		)
	}

	if !allowedSorts[query.SortBy] {

		return customErr.New(
			customErr.CodeInvalidSortField,
			"invalid sort field",
			http.StatusBadRequest,
			nil,
		)

	}

	if query.OrderBy != "asc" && query.OrderBy != "desc" {

		return customErr.New(
			customErr.CodeInvalidSortOrder,
			"sort order must be asc or desc",
			http.StatusBadRequest,
			nil,
		)
	}

	if query.OrderBy == "asc" {
		query.OrderBy = "ASC"
	}
	if query.OrderBy == "desc" {
		query.OrderBy = "DESC"
	}

	return nil

}
