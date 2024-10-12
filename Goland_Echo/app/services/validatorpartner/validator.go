package validatorpartner

import (
	"fmt"
	"reflect"

	"aprendiendoGo/app/services/response"

	v "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Validator es responsable de validar datos
type Validator struct {
	validate *v.Validate
	data     interface{}
}

// New crea y devuelve una instancia de Validator
func New() *Validator {
	return &Validator{
		validate: v.New(),
	}
}

// Data agrega un campo a validar
func (validator *Validator) Data(data interface{}) *Validator {
	validator.data = data
	return validator
}

// Validate valida todos los datos del validador
func (validator *Validator) Validate() map[string][]map[string]string {
	if err := validator.validate.Struct(validator.data); err != nil {
		return validator.makeDataErrsResponse(err)
	}
	return nil
}

// ValidateSend valida todos los datos del validador y responde
func (validator *Validator) ValidateSend(c echo.Context) error {
	if err := validator.validate.Struct(validator.data); err != nil {
		dataErrsResponse := validator.makeDataErrsResponse(err)
		return response.Unprocessable(c, "errors", dataErrsResponse)
	}
	return nil
}

// makeDataErrsResponse genera la respuesta de errores de validación personalizados
func (validator *Validator) makeDataErrsResponse(dataErrs interface{}) map[string][]map[string]string {
	dataErrsResponse := make(map[string][]map[string]string)

	for _, err := range dataErrs.(v.ValidationErrors) {
		errItem := make(map[string]string)
		field, _ := reflect.TypeOf(validator.data).FieldByName(err.Field())
		fieldTag, _ := field.Tag.Lookup("field")

		// Mensajes personalizados o genéricos
		key := fmt.Sprintf("%s.%s", err.Field(), err.Tag())
		if customMessage, exists := customErrorMessages[key]; exists {
			errItem["message"] = customMessage
		} else {
			errItem["message"] = fmt.Sprintf("%s no es válido", fieldTag)
		}

		errItem["field"] = fieldTag
		errItem["condition"] = err.ActualTag()
		dataErrsResponse["errors"] = append(dataErrsResponse["errors"], errItem)
	}

	return dataErrsResponse
}

// CustomValidation agrega una función para validar
func (validator *Validator) CustomValidation(tag string, fn func(fl v.FieldLevel) bool) *Validator {
	validator.validate.RegisterValidation(tag, fn)
	return validator
}
