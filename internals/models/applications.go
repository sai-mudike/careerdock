package models

import "time"

type Application struct {
	Id        string    `json:"id"`
	UserID    string    `json:"user_id"`
	JobID     string    `json:"job_id"`
	ResumeID  string    `json:"resume_id"`
	Status    string    `json:"status"`
	Notes     string    `json:"notes"`
	AppliedAT time.Time `json:"applied_at"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}

type ApplicationRequest struct {
	JobID     string    `json:"job_id" binding:"required"`
	ResumeID  string    `json:"resume_id"`
	Notes     string    `json:"notes" binding:"omitempty,max=500"`
	AppliedAT time.Time `json:"applied_at"`
}
type UpdateApplicationRequest struct {
	ResumeID  *string   `json:"resume_id"`
	Status    *string   `josn:"status" binding:"required"`
	Notes     *string   `json:"notes" binding:"omitempty,max=500"`
	AppliedAT time.Time `json:"applied_at"`
}

type ApplicationResponse struct {
	Id        string    `json:"id"`
	JobID     string    `json:"job_id"`
	ResumeID  string    `json:"resume_id"`
	Status    string    `json:"status"`
	Notes     string    `json:"notes"`
	AppliedAT time.Time `json:"applied_at"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}

func NewApplication(userID, jobId, resumeID, notes string) *Application {

	return &Application{
		UserID:   userID,
		JobID:    jobId,
		ResumeID: resumeID,
		Notes:    notes,
	}
}

func NewApplicationResponse(id, jobId, resumeID, status, notes string, appliedAt, createdAt, updatedAt time.Time) *ApplicationResponse {

	return &ApplicationResponse{
		Id:        id,
		JobID:     jobId,
		ResumeID:  resumeID,
		Status:    status,
		Notes:     notes,
		AppliedAT: appliedAt,
		CreatedAT: createdAt,
		UpdatedAT: updatedAt,
	}
}
