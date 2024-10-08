package validatorpartner

import (
	"fmt"
	"reflect"

	v "github.com/go-playground/validator/v10"
)

// Validator is in charge of validating data
type Validator struct {
	validate *v.Validate
	data     interface{}
}

// New Create and return an instance of Validator
func New() *Validator {

	var validator Validator

	validator.validate = v.New()

	return &validator
}

// Data add a field to validate
func (validator *Validator) Data(data interface{}) *Validator {

	validator.data = data

	return validator
}

// Validate validates all the validator data
func (validator *Validator) Validate() map[string][]map[string]string {

	dataErrs := validator.validate.Struct(validator.data)
	if dataErrs != nil {

		dataErrsResponse := validator.makeDataErrsResponse(dataErrs)

		return dataErrsResponse
	}

	return nil
}

// ValidateSend validates all the validator data and responde
// func (validator *Validator) ValidateSend(response *response.Response)  error {

// 	dataErrs := validator.validate.Struct(validator.data)
// 	if dataErrs != nil{

// 		dataErrsResponse := validator.makeDataErrsResponse(dataErrs);

// 		//defer panic("validator response")
// 		return response.Code(422).Data(dataErrsResponse).Message("errors").Send()

// 	}

// 	return nil
// }

// makeDataErrsResponse create the map with the errors
func (validator *Validator) makeDataErrsResponse(dataErrs interface{}) map[string][]map[string]string {

	dataErrsResponse := make(map[string][]map[string]string)

	for _, err := range dataErrs.(v.ValidationErrors) {
		errItem := make(map[string]string)
		field, _ := reflect.TypeOf(validator.data).FieldByName(err.Field())
		fieldTag, _ := field.Tag.Lookup("field")
		messageTag, _ := field.Tag.Lookup(err.ActualTag() + "-message")
		errItem["field"] = fieldTag
		errItem["condition"] = err.ActualTag()
		if messageTag != "" {
			errItem["message"] = messageTag
		} else {
			errItem["message"] = fmt.Sprint(err)
		}
		dataErrsResponse["errors"] = append(dataErrsResponse["errors"], errItem)
	}

	return dataErrsResponse

}

// CustomValidation add a function to validate
func (validator *Validator) CustomValidation(tag string, fn func(fl v.FieldLevel) bool) *Validator {

	validator.validate.RegisterValidation(tag, fn)

	return validator
}
