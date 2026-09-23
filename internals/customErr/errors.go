package customErr

type CustomErr struct {
	Code       string
	Message    string
	StatusCode int
	Err        error
}

func (e *CustomErr) Error() string {
	return e.Message
}

func (e *CustomErr) Unwrap() error {
	return e.Err
}

func New(code, message string, statusCode int, err error) *CustomErr {
	return &CustomErr{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
	}
}

const (
	// General
	CodeInvalidRequest      = "INVALID_REQUEST"
	CodeInvalidUUID         = "INVALID_UUID"
	CodeInternalServerError = "INTERNAL_SERVER_ERROR"

	// Authentication
	CodeUnauthorized       = "UNAUTHORIZED"
	CodeInvalidToken       = "INVALID_TOKEN"
	CodeExpiredToken       = "EXPIRED_TOKEN"
	CodeInvalidCredentials = "INVALID_CREDENTIALS"

	// User
	CodeUserNotFound       = "USER_NOT_FOUND"
	CodeEmailAlreadyExists = "EMAIL_ALREADY_EXISTS"

	// Jobs
	CodeJobNotFound           = "JOB_NOT_FOUND"
	CodeInvalidJobData        = "INVALID_JOB_DATA"
	CodeInvalidJobURL         = "INVALID_JOB_URL"
	CodeInvalidEmploymentType = "INVALID_EMPLOYMENT_TYPE"
	CodeInvalidPagination     = "INVALID_PAGINATION"
	CodeInvalidSortField      = "INVALID_SORT_FIELD"
	CodeInvalidSortOrder      = "INVALID_SORT_ORDER"

	// Resumes
	CodeResumeNotFound    = "RESUME_NOT_FOUND"
	CodeInvalidResumeName = "INVALID_RESUME_NAME"
	CodeInvalidFileType   = "INVALID_FILE_TYPE"
	CodeFileTooLarge      = "FILE_TOO_LARGE"
	CodeFileRequired      = "FILE_REQUIRED"
)
