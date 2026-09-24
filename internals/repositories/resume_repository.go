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

func CreateResume(ctx context.Context, resume models.Resume) (models.Resume, error) {
	query := `INSERT INTO resumes(user_id,name,file_name,file_path)
				VALUES($1,$2,$3,$4)
				RETURNING *;
`
	row := db.DB.QueryRowContext(ctx, query, resume.UserID, resume.Name, resume.FileName, resume.FilePath)

	if err := row.Err(); err != nil {
		return models.Resume{}, fmt.Errorf("create resume: %w", err)
	}

	var resumeFromDB models.Resume

	err := row.Scan(&resumeFromDB.Id, &resumeFromDB.UserID, &resumeFromDB.Name, &resumeFromDB.FileName, &resumeFromDB.FilePath, &resumeFromDB.CreatedAT)

	if err != nil {
		return models.Resume{}, fmt.Errorf("create resume: %w", err)
	}

	return resumeFromDB, nil
}

func GetResumes(ctx context.Context, userId string) ([]models.ResumeResponse, error) {

	query := `SELECT id,name,file_name,file_path,created_at FROM resumes WHERE user_id=$1;`

	rows, err := db.DB.QueryContext(ctx, query, userId)

	if err != nil {
		return nil, fmt.Errorf("Get all resume: %w", err)
	}

	defer rows.Close()

	var resumesFromDB = make([]models.ResumeResponse, 0)

	for rows.Next() {

		var singleResume models.ResumeResponse

		err := rows.Scan(&singleResume.Id, &singleResume.Name, &singleResume.FileName, &singleResume.FilePath, &singleResume.CreatedAT)

		if err != nil {
			return nil, fmt.Errorf("Get all resumes: Rows.Scan:  %w", err)
		}

		resumesFromDB = append(resumesFromDB, singleResume)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Get all resumes: %w", err)
	}

	return resumesFromDB, nil

}

func GetResumeByID(ctx context.Context, resumeID, userID string) (models.ResumeResponse, error) {
	query := `SELECT id,name,file_name,file_path,created_at FROM resumes WHERE id=$1 AND user_id=$2;`

	row := db.DB.QueryRowContext(ctx, query, resumeID, userID)

	var resumeFromDB models.ResumeResponse

	err := row.Scan(&resumeFromDB.Id, &resumeFromDB.Name, &resumeFromDB.FileName, &resumeFromDB.FilePath, &resumeFromDB.CreatedAT)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ResumeResponse{}, customErr.New(customErr.CodeResumeNotFound, "resume not found", http.StatusNotFound, err)
		}

		return models.ResumeResponse{}, fmt.Errorf("get resume by id: %w", err)
	}

	return resumeFromDB, nil
}

func DeleteResume(ctx context.Context, resumeId, userID string) error {
	query := `DELETE FROM resumes WHERE id=$1 AND user_id=$2;`

	_, err := db.DB.ExecContext(ctx, query, resumeId, userID)

	if err != nil {
		return fmt.Errorf("delete resume: %w", err)
	}

	return nil
}
