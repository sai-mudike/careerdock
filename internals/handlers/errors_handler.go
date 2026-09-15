package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	switch {

	// 400 Bad Request
	case errors.Is(err, customErr.ErrInvalidRequest):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_REQUEST",
			Message: fmt.Sprintf("Invalid request %v", err.Error()),
		})

	case errors.Is(err, customErr.ErrInvalidJobData):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_JOB_DATA",
			Message: "Invalid job data",
		})

	case errors.Is(err, customErr.ErrInvalidSalaryRange):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_SALARY_RANGE",
			Message: "Minimum salary cannot be greater than maximum salary",
		})

	case errors.Is(err, customErr.ErrInvalidStatus):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_STATUS",
			Message: "Invalid application status",
		})

	case errors.Is(err, customErr.ErrInvalidJob_url):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_JOB_URL",
			Message: "Invalid Job url "})

	case errors.Is(err, customErr.ErrInvalidEmployment_type):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_EMPLOYMNET_TYPE",
			Message: "Invalid employment type"})

	case errors.Is(err, customErr.ErrInvalidApplied_at):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_APPLIED_AT",
			Message: "Applied at Cannot be after the application creating data"})
	case errors.Is(err, customErr.ErrInvalidPagination):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_PAGINATION",
			Message: "Valid page number is required"})
	case errors.Is(err, customErr.ErrInvalidSortOrder):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_SORTING_ORDER",
			Message: "sorting can be either 'asc' or 'desc'"})
	case errors.Is(err, customErr.ErrInvalidSort):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_SORTING",
			Message: "invalid sorting data"})

	case errors.Is(err, customErr.ErrInvalidResumeData):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_RESUME_DATA",
			Message: "invalid resume data"})

	case errors.Is(err, customErr.ErrInvalidResumeFile_path):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_RESUME_FILE_PATH",
			Message: "invalid resume file path"})

	// 401 Unauthorized
	case errors.Is(err, customErr.ErrMissingToken):
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    "MISSING_TOKEN",
			Message: "Authentication token is required",
		})

	case errors.Is(err, customErr.ErrInvalidToken):
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    "INVALID_TOKEN",
			Message: "Invalid authentication token",
		})

	case errors.Is(err, customErr.ErrTokenExpired):
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    "TOKEN_EXPIRED",
			Message: "Authentication token has expired",
		})

	case errors.Is(err, customErr.ErrInvalidCredentials):
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    "INVALID_CREDENTIALS",
			Message: "Invalid email or password",
		})

	// 403 Forbidden
	case errors.Is(err, customErr.ErrForbidden):
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			Code:    "FORBIDDEN",
			Message: "You do not have permission to perform this action",
		})

	// 404 Not Found
	case errors.Is(err, customErr.ErrUserNotFound):
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    "USER_NOT_FOUND",
			Message: "User not found",
		})

	case errors.Is(err, customErr.ErrJobNotFound):
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    "JOB_NOT_FOUND",
			Message: "job not found",
		})
	case errors.Is(err, customErr.ErrResumeNotFound):
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    "RESUME_NOT_FOUND",
			Message: "resume not found",
		})
	case errors.Is(err, customErr.ErrApplicationJobOrResumeNotFound):
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    "JOB_OR_RESUME_NOT_FOUND",
			Message: "job or resume not found or does not belong to you",
		})
	// 409 Conflict
	case errors.Is(err, customErr.ErrEmailAlreadyExists):
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Code:    "EMAIL_ALREADY_EXISTS",
			Message: "An account with this email already exists",
		})

	case errors.Is(err, customErr.ErrDuplicateJob):
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Code:    "DUPLICATE_JOB",
			Message: "You already have this job tracked",
		})
	case errors.Is(err, customErr.ErrResumeAlreadyExists):
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Code:    "RESUME_ALREADY_EXISTS",
			Message: "An account with this resume already exists",
		})

		// 500 Internal Server Error

	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: fmt.Sprintf("An unexpected error occurred %v", err.Error()),
		})

	}

}
