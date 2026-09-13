package models

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	Id             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
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
