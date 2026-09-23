package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/db"
	"github.com/sai-mudike/careerdock.git/internals/models"
)

func CreateJob(ctx context.Context, job models.Job) (models.Job, error) {
	query := `
INSERT INTO jobs(user_id,company_name,position,job_url,location,employment_type,salary_min,salary_max,description)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
RETURNING *;
`
	smt, err := db.DB.Prepare(query)

	if err != nil {
		return models.Job{}, fmt.Errorf("Create job: %w", err)
	}

	defer smt.Close()

	row := smt.QueryRowContext(ctx, job.UserID, job.CompanyName, job.Position, job.JobURL, job.Location, job.EmploymentType, job.SalaryMin, job.SalaryMax, job.Description)

	var jobFromDB models.Job

	err = row.Scan(&jobFromDB.Id, &jobFromDB.UserID, &jobFromDB.CompanyName, &jobFromDB.Position, &jobFromDB.JobURL, &jobFromDB.Location, &jobFromDB.EmploymentType, &jobFromDB.SalaryMin, &jobFromDB.SalaryMax, &jobFromDB.Description, &jobFromDB.CreatedAT, &jobFromDB.UpdatedAT)

	if err != nil {
		return models.Job{}, fmt.Errorf("Create job: Row.Scan: %w", err)
	}

	return jobFromDB, nil

}

func GetAllJobs(ctx context.Context, userId string, jobQuery models.JobQuery) ([]models.JobResponse, error) {
	query := `
	SELECT id,company_name,position,job_url,location,employment_type,salary_min,salary_max,description,created_at,updated_at FROM jobs WHERE user_id=$1
	`

	queryArgs := []any{userId}
	argsCount := 2

	if jobQuery.OrderBy != "" {
		query += fmt.Sprintf(" ORDER BY %s %s", jobQuery.SortBy, jobQuery.OrderBy)
	}

	if jobQuery.Page >= 1 {

		offset := (jobQuery.Page - 1) * jobQuery.Limit
		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argsCount, argsCount+1)
		queryArgs = append(queryArgs, jobQuery.Limit, offset)
	}

	rows, err := db.DB.QueryContext(ctx, query, queryArgs...)

	if err != nil {
		return nil, fmt.Errorf("Get jobs: %w", err)
	}

	defer rows.Close()

	var jobsList = make([]models.JobResponse, 0)
	for rows.Next() {
		var singleJob models.JobResponse

		err := rows.Scan(&singleJob.Id, &singleJob.CompanyName, &singleJob.Position, &singleJob.JobURL, &singleJob.Location, &singleJob.EmploymentType, &singleJob.SalaryMin, &singleJob.SalaryMax, &singleJob.Description, &singleJob.CreatedAT, &singleJob.UpdatedAT)

		if err != nil {
			return nil, fmt.Errorf("Get jobs: Rows.Scan: %w", err)
		}

		jobsList = append(jobsList, singleJob)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Get jobs: %w", err)
	}

	return jobsList, nil
}

func GetJobByID(ctx context.Context, jobID string, userID string) (models.JobResponse, error) {
	query := `
SELECT id,company_name,position,job_url,location,employment_type,salary_min,salary_max,description,created_at,updated_at FROM jobs WHERE id=$1 AND user_id=$2;
`

	row := db.DB.QueryRowContext(ctx, query, jobID, userID)

	var singleJob models.JobResponse

	err := row.Scan(&singleJob.Id, &singleJob.CompanyName, &singleJob.Position, &singleJob.JobURL, &singleJob.Location, &singleJob.EmploymentType, &singleJob.SalaryMin, &singleJob.SalaryMax, &singleJob.Description, &singleJob.CreatedAT, &singleJob.UpdatedAT)

	if errors.Is(err, sql.ErrNoRows) {
		return models.JobResponse{}, customErr.New(customErr.CodeJobNotFound, "job not found", http.StatusNotFound, err)
	}

	if err != nil {
		return models.JobResponse{}, fmt.Errorf("Get job by id: %w", err)
	}

	return singleJob, nil
}

func UpdateJob(ctx context.Context, jobID string, job models.Job) (models.Job, error) {

	query := `
	UPDATE jobs
	SET company_name=$3,position=$4,job_url=$5,location=$6,employment_type=$7,salary_min=$8,salary_max=$9,description=$10,updated_at=CURRENT_TIMESTAMP
	WHERE id=$1 AND user_id=$2
	RETURNING *;
	`

	smt, err := db.DB.PrepareContext(ctx, query)
	if err != nil {
		return models.Job{}, fmt.Errorf("Update job: %w", err)
	}

	defer smt.Close()

	row := smt.QueryRowContext(ctx, jobID, job.UserID, job.CompanyName, job.Position, job.JobURL, job.Location, job.EmploymentType, job.SalaryMin, job.SalaryMax, job.Description)

	if err := row.Err(); err != nil {
		return models.Job{}, fmt.Errorf("update job: %w", err)
	}

	var jobFromDB models.Job

	err = row.Scan(&jobFromDB.Id, &jobFromDB.UserID, &jobFromDB.CompanyName, &jobFromDB.Position, &jobFromDB.JobURL, &jobFromDB.Location, &jobFromDB.EmploymentType, &jobFromDB.SalaryMin, &jobFromDB.SalaryMax, &jobFromDB.Description, &jobFromDB.CreatedAT, &jobFromDB.UpdatedAT)

	if err != nil {
		return models.Job{}, fmt.Errorf("Update job: Rows.Scan: %w", err)
	}

	return jobFromDB, nil

}

func DeleteJob(ctx context.Context, JobID, userId string) error {
	query := `
	DELETE FROM jobs WHERE id=$1 AND user_id=$2;
	`

	_, err := db.DB.ExecContext(ctx, query, JobID, userId)

	if err != nil {

		return fmt.Errorf("Delete job: Rows.Scan: %w", err)
	}
	return nil

}
