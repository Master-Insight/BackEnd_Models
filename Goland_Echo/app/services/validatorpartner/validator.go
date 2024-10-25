package validatorpartner

import (
	"fmt"
	"reflect"

	"aprendiendoGo/app/services/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Validator es responsable de validar datos
type Validator struct {
	validate *validator.Validate
	data     interface{}
}

// New crea y devuelve una instancia de Validator
func New() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

// Data agrega los datos que se van a validar
func (v *Validator) Data(data interface{}) *Validator {
	v.data = data
	return v
}

// Validate realiza la validación de los datos y devuelve un mapa con los errores
func (v *Validator) Validate() map[string][]map[string]string {
	if err := v.validate.Struct(v.data); err != nil {
		return v.makeErrorResponse(err.(validator.ValidationErrors)) // Cambio aquí para castear correctamente
	}
	return nil
}

// ValidateSend valida los datos y envía la respuesta en caso de errores
func (v *Validator) ValidateSend(c echo.Context) error {
	if err := v.validate.Struct(v.data); err != nil {
		errorResponse := v.makeErrorResponse(err.(validator.ValidationErrors))

		//return response.Code(422).Data(errorResponse).Message("errors").Send()
		return response.Unprocessable(c, "validation_errors", errorResponse)
	}
	return nil // Retorna nil si no hay errores
}

// makeErrorResponse crea un mapa con los mensajes de error personalizados
func (v *Validator) makeErrorResponse(validationErrors validator.ValidationErrors) map[string][]map[string]string {
	errors := make(map[string][]map[string]string)

	for _, err := range validationErrors {
		field, _ := reflect.TypeOf(v.data).FieldByName(err.Field())
		fieldTag := field.Tag.Get("field")
		errorMessage := v.getCustomMessage(err, fieldTag)

		errors["errors"] = append(errors["errors"], map[string]string{
			"field":     fieldTag,
			"condition": err.Tag(),
			"message":   errorMessage,
		})
	}
	return errors
}

// getCustomMessage devuelve el mensaje personalizado o uno genérico si no existe
func (v *Validator) getCustomMessage(err validator.FieldError, fieldTag string) string {
	key := fmt.Sprintf("%s.%s", err.Field(), err.Tag())
	if customMsg, exists := customErrorMessages[key]; exists {
		return customMsg
	}
	return fmt.Sprintf("%s no es válido", fieldTag)
}

// CustomValidation agrega una función de validación personalizada
func (v *Validator) CustomValidation(tag string, fn validator.Func) *Validator {
	v.validate.RegisterValidation(tag, fn)
	return v
}
