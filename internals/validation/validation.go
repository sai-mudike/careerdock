package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationMessage(err error) string {
	var validationErrors validator.ValidationErrors

	if !errors.As(err, &validationErrors) {
		return "invalid request"
	}

	fieldErr := validationErrors[0]

	switch fieldErr.Tag() {
	case "required":
		return fmt.Sprintf(
			"%s is required",
			fieldErr.Field(),
		)

	case "min":
		return fmt.Sprintf(
			"%s must be at least %s characters",
			fieldErr.Field(),
			fieldErr.Param(),
		)

	case "max":
		return fmt.Sprintf(
			"%s must not exceed %s characters",
			fieldErr.Field(),
			fieldErr.Param(),
		)

	case "email":
		return fmt.Sprintf(
			"%s must be a valid email",
			fieldErr.Field(),
		)

	default:
		return fmt.Sprintf(
			"%s is invalid",
			fieldErr.Field(),
		)
	}
}
