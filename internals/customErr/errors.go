package customErr

import "errors"

var (
	// 400 Bad Request
	ErrInvalidRequest         = errors.New("invalid request")
	ErrInvalidJobData         = errors.New("invalid job data")
	ErrInvalidSalaryRange     = errors.New("invalid salary range")
	ErrInvalidStatus          = errors.New("invalid job status")
	ErrInvalidJob_url         = errors.New("invalid job url")
	ErrInvalidEmployment_type = errors.New("invalid employment type")
	ErrInvalidApplied_at      = errors.New("invalid Applied at")
	ErrInvalidPagination      = errors.New("invalid page number")
	ErrInvalidSort            = errors.New("invalid sorting value")
	ErrInvalidSortOrder       = errors.New("invalid sorting order")
	ErrInvalidResumeData      = errors.New("invalid Resume data")
	ErrInvalidResumeFile_path = errors.New("invalid resume file path")

	// 401 Unauthorized
	ErrMissingToken       = errors.New("missing token")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("Invalid or malformed authentication")

	// 403 Forbidden
	ErrForbidden = errors.New("forbidden")

	// 404 Not Found
	ErrUserNotFound                   = errors.New("user not found")
	ErrJobNotFound                    = errors.New("job not found")
	ErrResumeNotFound                 = errors.New("Resume not found")
	ErrApplicationAlreadyExists       = errors.New("Application not found")
	ErrApplicationJobOrResumeNotFound = errors.New("job or resume not found")

	// 409 Conflict
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrDuplicateJob        = errors.New("job already exists")
	ErrResumeAlreadyExists = errors.New("resume already exists")

	// 500 Internal Server Error
	ErrInternal        = errors.New("internal server error")
	ErrUserNotCreated  = errors.New("unable to register user.")
	ErrTokenGeneration = errors.New("Failed to issue session token. Please try again later.")
)
