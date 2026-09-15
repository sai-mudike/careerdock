package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/db"
	"github.com/sai-mudike/careerdock.git/internals/models"
)

func CreateJob(ctx context.Context, job models.Job) (models.Job, error) {
	query := `
INSERT INTO jobs(user_id,company_name,position,job_url,location,employment_type,salary_min,salary_max,description)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
RETURNING id,user_id,company_name,position,job_url,location,employment_type,salary_min,salary_max,description;
`
	smt, err := db.DB.Prepare(query)

	if err != nil {
		return models.Job{}, err
	}

	defer smt.Close()

	row := smt.QueryRowContext(ctx, job.UserID, job.CompanyName, job.Position, job.JobURL, job.Location, job.EmploymentType, job.SalaryMin, job.SalaryMax, job.Description)

	var jobFromDB models.Job

	err = row.Scan(&jobFromDB.Id, &jobFromDB.UserID, &jobFromDB.CompanyName, &jobFromDB.Position, &jobFromDB.JobURL, &jobFromDB.Location, &jobFromDB.EmploymentType, &jobFromDB.SalaryMin, &jobFromDB.SalaryMax, &jobFromDB.Description)

	if err != nil {
		return models.Job{}, customErr.ErrInternal
	}

	return jobFromDB, nil

}

func GetAllJobs(ctx context.Context) ([]models.Job, error) {
	query := `
	SELECT * FROM jobs;
	`
	rows, err := db.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobsList []models.Job
	for rows.Next() {
		var singleJob models.Job

		err := rows.Scan(&singleJob.Id, &singleJob.UserID, &singleJob.CompanyName, &singleJob.Position, &singleJob.JobURL, &singleJob.Location, &singleJob.EmploymentType, &singleJob.SalaryMin, &singleJob.SalaryMax, &singleJob.Description, &singleJob.CreatedAT, &singleJob.UpdatedAT)

		if err != nil {
			return nil, customErr.ErrInternal
		}

		jobsList = append(jobsList, singleJob)
	}

	if err := rows.Err(); err != nil {
		return nil, customErr.ErrInternal
	}

	return jobsList, nil
}

func GetJobByID(ctx context.Context, jobID uuid.UUID) (models.Job, error) {
	query := `
SELECT * FROM jobs WHERE id=$1;
`

	row := db.DB.QueryRowContext(ctx, query, jobID)

	if err := row.Err(); err != nil {
		return models.Job{}, err
	}

	var singleJob models.Job

	err := row.Scan(&singleJob.Id, &singleJob.UserID, &singleJob.CompanyName, &singleJob.Position, &singleJob.JobURL, &singleJob.Location, &singleJob.EmploymentType, &singleJob.SalaryMin, &singleJob.SalaryMax, &singleJob.Description, &singleJob.CreatedAT, &singleJob.UpdatedAT)

	if err != nil {
		return models.Job{}, customErr.ErrJobNotFound
	}

	return singleJob, nil
}

func UpdateJob(ctx context.Context, job models.Job) (models.Job, error) {

	query := `
	UPDATE jobs
	SET company_name=$2,position=$3,job_url=$4,location=$5,employment_type=$6,salary_min=$7,salary_max=$8,description=$9,updated_at=CURRENT_TIMESTAMP
	WHERE id=$1
	RETURNING *;
	`

	smt, err := db.DB.PrepareContext(ctx, query)
	if err != nil {
		return models.Job{}, err
	}

	defer smt.Close()

	row := smt.QueryRowContext(ctx, job.Id, job.CompanyName, job.Position, job.JobURL, job.Location, job.EmploymentType, job.SalaryMin, job.SalaryMax, job.Description)

	if err := row.Err(); err != nil {
		return models.Job{}, customErr.ErrInternal
	}

	var singleJob models.Job

	err = row.Scan(&singleJob.Id, &singleJob.UserID, &singleJob.CompanyName, &singleJob.Position, &singleJob.JobURL, &singleJob.Location, &singleJob.EmploymentType, &singleJob.SalaryMin, &singleJob.SalaryMax, &singleJob.Description, &singleJob.CreatedAT, &singleJob.UpdatedAT)

	if err != nil {
		return models.Job{}, customErr.ErrInternal
	}

	return singleJob, nil

}

func DeleteJob(ctx context.Context, JobID uuid.UUID) error {
	query := `
	DELETE FROM jobs WHERE id=$1;
	`

	_, err := db.DB.ExecContext(ctx, query, JobID)

	if err != nil {

		return customErr.ErrInternal
	}
	return nil

}
