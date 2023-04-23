package controllers

import "github.com/go-playground/validator"

type ErrorResponse struct {
	FailedField string
	Tag         string
}

var validate = validator.New()

//Validate
func ValidateStruct(req interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
	}
	return errors
}
