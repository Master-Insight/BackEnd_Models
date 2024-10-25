package errs

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Error es una implementación personalizada del error que incluye lógica de respuesta
type Error struct {
	Message      string `json:"message"`
	Description  string `json:"description,omitempty"`
	CodeHTTP     int    `json:"code_http"`
	CodeNameHTTP string `json:"code_name_http"`
	ErrorWrapped error  `json:"error_wrapped,omitempty"`
}

// NewError crea un nuevo error personalizado
func New(message string, description string, codeHTTP int) *Error {
	return &Error{
		Message:      message,
		Description:  description,
		CodeHTTP:     codeHTTP,
		CodeNameHTTP: http.StatusText(codeHTTP),
	}
}

// Error convierte el ErrorResponse en string
func (e *Error) Error() string {
	return e.Message
}

// SetWrapped añade un error interno para facilitar el rastreo
func (e *Error) SetWrapped(err error) *Error {
	e.ErrorWrapped = err
	return e
}

// FullError devuelve el error completo en formato JSON
func (e *Error) FullError() string {
	fullError, err := json.Marshal(e)
	if err != nil {
		return e.Error()
	}
	return string(fullError)
}

// Send convierte el error en una respuesta HTTP
func (e *Error) Send(c echo.Context) error {
	return c.JSON(e.CodeHTTP, map[string]interface{}{
		"message":        e.Message,
		"description":    e.Description,
		"code_http":      e.CodeHTTP,
		"code_name_http": e.CodeNameHTTP,
	})
}
