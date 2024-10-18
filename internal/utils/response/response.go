package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func WriteResponse(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// Custom error message map for each validation error
func validationErrorMessages(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s", err.Field(), err.Param()) // For string length, number, etc.
	case "max":
		return fmt.Sprintf("%s must be at most %s", err.Field(), err.Param()) // Max length, number, etc.
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters long", err.Field(), err.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())
	case "url":
		return fmt.Sprintf("%s must be a valid URL", err.Field())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", err.Field(), err.Param()) // For numbers, string length, etc.
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", err.Field(), err.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", err.Field(), err.Param()) // For numbers, string length, etc.
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", err.Field(), err.Param())
	case "eq":
		return fmt.Sprintf("%s must be equal to %s", err.Field(), err.Param())
	case "ne":
		return fmt.Sprintf("%s must not be equal to %s", err.Field(), err.Param())
	case "alpha":
		return fmt.Sprintf("%s must contain only alphabetic characters", err.Field())
	case "alphanum":
		return fmt.Sprintf("%s must contain only alphanumeric characters", err.Field())
	case "numeric":
		return fmt.Sprintf("%s must be a valid number", err.Field())
	case "uuid":
		return fmt.Sprintf("%s must be a valid UUID", err.Field())
	case "uuid4":
		return fmt.Sprintf("%s must be a valid UUID version 4", err.Field())
	case "boolean":
		return fmt.Sprintf("%s must be a valid boolean", err.Field())
	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", err.Field(), err.Param()) // E.g., "male female" for oneof tag
	case "contains":
		return fmt.Sprintf("%s must contain the value %s", err.Field(), err.Param())
	case "containsany":
		return fmt.Sprintf("%s must contain at least one of the characters %s", err.Field(), err.Param())
	case "excludes":
		return fmt.Sprintf("%s must not contain the value %s", err.Field(), err.Param())
	case "excludesall":
		return fmt.Sprintf("%s must not contain any of the characters %s", err.Field(), err.Param())
	case "startswith":
		return fmt.Sprintf("%s must start with %s", err.Field(), err.Param())
	case "endswith":
		return fmt.Sprintf("%s must end with %s", err.Field(), err.Param())
	case "ip":
		return fmt.Sprintf("%s must be a valid IP address", err.Field())
	case "ipv4":
		return fmt.Sprintf("%s must be a valid IPv4 address", err.Field())
	case "ipv6":
		return fmt.Sprintf("%s must be a valid IPv6 address", err.Field())
	case "mac":
		return fmt.Sprintf("%s must be a valid MAC address", err.Field())
	case "hexadecimal":
		return fmt.Sprintf("%s must be a valid hexadecimal", err.Field())
	case "base64":
		return fmt.Sprintf("%s must be a valid base64 string", err.Field())
	case "datetime":
		return fmt.Sprintf("%s must be a valid datetime in the format %s", err.Field(), err.Param())
	case "file":
		return fmt.Sprintf("%s must be a valid file", err.Field()) // Custom validation (if you add it)
	case "struct":
		return fmt.Sprintf("%s must be a valid struct", err.Field())
	case "slice":
		return fmt.Sprintf("%s must be a valid slice", err.Field())
	case "array":
		return fmt.Sprintf("%s must be a valid array", err.Field())
	case "map":
		return fmt.Sprintf("%s must be a valid map", err.Field())
	case "string":
		return fmt.Sprintf("%s must be a valid string", err.Field())
	case "int":
		return fmt.Sprintf("%s must be a valid integer", err.Field())
	case "float":
		return fmt.Sprintf("%s must be a valid float", err.Field())
	case "bool":
		return fmt.Sprintf("%s must be a valid boolean", err.Field())

	default:
		return fmt.Sprintf("%s is not valid", err.Field())
	}
}

func ValidateStruct(data interface{}) map[string]string {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		fieldName := strings.ToLower(err.Field())
		errors[fieldName] = validationErrorMessages(err)
	}
	return errors
}
