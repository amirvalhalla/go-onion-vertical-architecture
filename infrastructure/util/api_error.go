package util

import "github.com/go-playground/validator/v10"

func GetErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "field is required"
	case "uuid4":
		return "field has invalid uuid"
	case "min":
		return "field should be greater than " + fe.Param()
	case "max":
		return "field should be less than " + fe.Param()
	case "uppercase":
		return "field should be uppercase"
	case "lte":
		return "field should be less than " + fe.Param()
	case "gte":
		return "field should be greater than " + fe.Param()
	}
	return "unknown error"
}
