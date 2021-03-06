package validation

import "errors"

func ValidateNil(name string, value interface{}) []error {
	var result []error = []error{}

	if value == nil {
		result = append(result, errors.New("Field "+name+" cannot be nill"))
	}

	return result
}

func ValidateString(name string, value string) []error {
	var result []error = ValidateNil(name, value)

	if len(result) == 0 && len(value) == 0 {
		result = append(result, errors.New("Field "+name+" cannot be empty"))
	}

	return result
}
