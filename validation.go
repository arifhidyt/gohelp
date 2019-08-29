package gohelp

import (
	"fmt"
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

/**
 * DoValidate is process validation
 * @return object
 */
func DoValidate(input interface{}) map[string]interface{} {

	validate = validator.New()
	var invalidParamsMessage []string

	validate.RegisterValidation("todate", validateDateTime)
	validate.RegisterValidation("ISO8601date", validateISO8601)

	if castedObject, ok := validate.Struct(input).(validator.ValidationErrors); ok {
		//var fields []string
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				invalidParamsMessage = append(invalidParamsMessage, fmt.Sprintf("%s is required",
					err.Field()))
			case "email":
				invalidParamsMessage = append(invalidParamsMessage, fmt.Sprintf("%s is not valid email",
					err.Field()))
			case "todate":
				invalidParamsMessage = append(invalidParamsMessage, fmt.Sprintf("%s is not valid date",
					err.Field()))
			case "ISO8601date":
				invalidParamsMessage = append(invalidParamsMessage, fmt.Sprintf("%s is not valid datetime",
					err.Field()))
			case "gte":
				invalidParamsMessage = append(invalidParamsMessage, fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param()))
			case "lte":
				invalidParamsMessage = append(invalidParamsMessage, fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param()))
			case "max":
				invalidParamsMessage = append(invalidParamsMessage, fmt.Sprintf("%s value Maximum character %s",
					err.Field(), err.Param()))
			case "min":
				invalidParamsMessage = append(invalidParamsMessage, fmt.Sprintf("%s value Minimum character %s",
					err.Field(), err.Param()))
			}
		}
	}

	if len(invalidParamsMessage) > 0 {
		return map[string]interface{}{
			"validateError": invalidParamsMessage,
		}
	}

	return nil
}

// validateDateTime implements validator.Func
// Date time format yyyy-mm-dd
func validateDateTime(fl validator.FieldLevel) bool {
	regexString := "^(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])$"
	regexResult := regexp.MustCompile(regexString)

	return regexResult.MatchString(fl.Field().String())
}

// Date time format ISO8601
func validateISO8601(fl validator.FieldLevel) bool {
	regexString := "^(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])(?:T|\\s)(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])?(Z)?$"
	regexResult := regexp.MustCompile(regexString)

	return regexResult.MatchString(fl.Field().String())
}
