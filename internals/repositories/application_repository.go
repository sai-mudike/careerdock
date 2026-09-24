package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/db"
	"github.com/sai-mudike/careerdock.git/internals/models"
)

func CreateApplication(ctx context.Context, application models.Application) (models.Application, error) {

	query := `INSERT INTO applications(user_id,job_id,resume_id,notes,applied_at) 
	VALUES ($1,$2,$3,$4,$5)
	RETURNING *`

	row := db.DB.QueryRowContext(ctx, query, application.UserID, application.JobID, application.ResumeID, application.Notes, application.AppliedAT)

	var applicationFromDB models.Application

	err := row.Scan(&applicationFromDB.Id, &applicationFromDB.UserID, &applicationFromDB.JobID, &applicationFromDB.ResumeID, &applicationFromDB.Status, &applicationFromDB.Notes, &applicationFromDB.AppliedAT, &applicationFromDB.CreatedAT, &applicationFromDB.UpdatedAT)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return models.Application{}, customErr.New(customErr.CodeApplicationAlreadyExists, "application already exists for this job", http.StatusConflict, err)
			}
		}

		return models.Application{}, fmt.Errorf("create Application: %v", err)
	}

	return applicationFromDB, nil

}

func GetAllApplications(ctx context.Context, userID string) ([]models.ApplicationResponse, error) {

	query := `SELECT id,job_id,resume_id,status,notes,applied_at,created_at,updated_at FROM applications WHERE user_id=$1`

	rows, err := db.DB.QueryContext(ctx, query, userID)

	if err != nil {
		return nil, fmt.Errorf("get all application: %v", err)
	}

	var applicationsFromDB = make([]models.ApplicationResponse, 0)

	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get all applications: interate rows: %v", err)
	}

	for rows.Next() {

		var singleApplication models.ApplicationResponse

		err := rows.Scan(&singleApplication.Id, &singleApplication.JobID, &singleApplication.ResumeID, &singleApplication.Status, &singleApplication.Notes, &singleApplication.AppliedAT, &singleApplication.CreatedAT, &singleApplication.UpdatedAT)

		if err != nil {
			return nil, fmt.Errorf("get All Applications: Rows.Scan: %v", err)
		}
		applicationsFromDB = append(applicationsFromDB, singleApplication)
	}

	return applicationsFromDB, nil

}

func GetApplicationByID(ctx context.Context, applicationID, userID string) (models.ApplicationResponse, error) {
	query := `SELECT id,job_id,resume_id,status,notes,applied_at,created_at,updated_at FROM applications WHERE id=$1 AND user_id=$2 `

	row := db.DB.QueryRowContext(ctx, query, applicationID, userID)

	var singleApplication models.ApplicationResponse

	err := row.Scan(&singleApplication.Id, &singleApplication.JobID, &singleApplication.ResumeID, &singleApplication.Status, &singleApplication.Notes, &singleApplication.AppliedAT, &singleApplication.CreatedAT, &singleApplication.UpdatedAT)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return models.ApplicationResponse{}, customErr.New(customErr.CodeApplicationNotFound, "application not found", http.StatusNotFound, err)
		}
		return models.ApplicationResponse{}, fmt.Errorf("get Application by id: Rows.Scan: %v", err)

	}

	return singleApplication, nil

}

func UpdateApplications(ctx context.Context, application models.Application) (models.Application, error) {

	query := `UPDATE applications
	SET resume_id=$3,status=$4,notes=$5,applied_at=$6,updated_at=CURRENT_TIMESTAMP
	WHERE id=$1 AND user_id=$2
	RETURNING *`

	row := db.DB.QueryRowContext(ctx, query, application.Id, application.UserID, application.ResumeID, application.Status, application.Notes, application.AppliedAT)

	var applicationFromDB models.Application

	err := row.Scan(&applicationFromDB.Id, &applicationFromDB.UserID, &applicationFromDB.JobID, &applicationFromDB.ResumeID, &applicationFromDB.Status, &applicationFromDB.Notes, &applicationFromDB.AppliedAT, &applicationFromDB.CreatedAT, &applicationFromDB.UpdatedAT)

	if err != nil {

		return models.Application{}, fmt.Errorf("update Application: %v", err)
	}

	return applicationFromDB, nil

}

func DeleteApplication(ctx context.Context, applicationId, userID string) error {

	query := `DELETE FROM applications WHERE id=$1 AND user_id=$2`

	_, err := db.DB.ExecContext(ctx, query, applicationId, userID)

	if err != nil {
		return fmt.Errorf("delete application: %v", err)
	}

	return nil

}
