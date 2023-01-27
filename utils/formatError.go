package utils

import (
	"errors"
	"fmt"

	"github.com/elisalimli/meetmeup/models"
	"github.com/go-playground/validator/v10"
)

func MsgForTag(fe validator.FieldError) string {
	fieldName := fe.StructField()
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%v is required", fieldName)
	case "email":
		return "Invalid email"
	case "max":
		return fmt.Sprintf("%v must be less than %v characters", fieldName, fe.Param())
	}
	return fe.Error() // default error
}

func FormatErrors(err error) []*models.FieldError {
	var ve validator.ValidationErrors
	fmt.Println("ve", &ve)
	if errors.As(err, &ve) {
		out := make([]*models.FieldError, len(ve))
		for i, fe := range ve {
			out[i] = &models.FieldError{Path: fe.Field(), Message: MsgForTag(fe)}
		}
		return out
	}
	return nil
}
