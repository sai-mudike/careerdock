package models

import (
	"time"
)

type Job struct {
	Id             string    `json:"id"`
	UserID         string    `json:"user_id"`
	CompanyName    string    `json:"company_name" binding:"required"`
	Position       string    `json:"position" binding:"required"`
	JobURL         string    `json:"job_url"`
	Location       string    `json:"location"`
	EmploymentType string    `json:"employment_type" binding:"required"`
	SalaryMin      int       `json:"salary_min"`
	SalaryMax      int       `json:"salary_max"`
	Description    string    `json:"description"`
	CreatedAT      time.Time `json:"created_at"`
	UpdatedAT      time.Time `json:"updated_at"`
}

type JobRequest struct {
	CompanyName    string `json:"company_name" binding:"required,min=2,max=250"`
	Position       string `json:"position" binding:"required,min=2,max=250"`
	JobURL         string `json:"job_url" binding:"url"`
	Location       string `json:"location" binding:"max=250"`
	EmploymentType string `json:"employment_type" binding:"required"`
	SalaryMin      int    `json:"salary_min" binding:"gte=0"`
	SalaryMax      int    `json:"salary_max"  binding:"gte=0"`
	Description    string `json:"description" binding:"max=500"`
}

type JobResponse struct {
	Id             string    `json:"id"`
	CompanyName    string    `json:"company_name"`
	Position       string    `json:"position"`
	JobURL         string    `json:"job_url"`
	Location       string    `json:"location"`
	EmploymentType string    `json:"employment_type"`
	SalaryMin      int       `json:"salary_min"`
	SalaryMax      int       `json:"salary_max"`
	Description    string    `json:"description"`
	CreatedAT      time.Time `json:"created_at"`
	UpdatedAT      time.Time `json:"updated_at"`
}

func NewJob(UserID string, CompanyName string, Position string, JobURL string, Location string, EmploymentType string, SalaryMin int, SalaryMax int, Description string) *Job {

	return &Job{
		UserID:         UserID,
		CompanyName:    CompanyName,
		Position:       Position,
		JobURL:         JobURL,
		Location:       Location,
		EmploymentType: EmploymentType,
		SalaryMin:      SalaryMin,
		SalaryMax:      SalaryMax,
		Description:    Description,
	}

}

func NewJobResponse(id string, CompanyName string, Position string, JobURL string, Location string, EmploymentType string, SalaryMin int, SalaryMax int, Description string, CreatedAt time.Time, UpdatedAt time.Time) *JobResponse {

	return &JobResponse{
		Id:             id,
		CompanyName:    CompanyName,
		Position:       Position,
		JobURL:         JobURL,
		Location:       Location,
		EmploymentType: EmploymentType,
		SalaryMin:      SalaryMin,
		SalaryMax:      SalaryMax,
		Description:    Description,
		CreatedAT:      CreatedAt,
		UpdatedAT:      UpdatedAt,
	}

}
