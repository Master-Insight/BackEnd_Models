package res

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//   ! MULTIPLES (Falta hacer)
//   res.sendSuccessOrNotFound = (variable, title = "Item") => (variable) ? res.sendSuccess(variable) : res.sendNotFound(`${title} not found`);
//   res.sendCatchError = (error, message = "Internal Server Error") => res.sendServerError(message, error.toString());

// Success envía una respuesta con código 200
func Success(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Success"
	}
	// Si no hay datos, se deja el valor por defecto (vacío)
	resp := New(c).Message(message).Code(http.StatusOK)
	if data != nil {
		resp.Data(data)
	}
	return resp.Send()
}

// Created envía una respuesta con código 201
func Created(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Created"
	}
	// Si no hay datos, se deja el valor por defecto (vacío)
	resp := New(c).Message(message).Code(http.StatusOK)
	if data != nil {
		resp.Data(data)
	}
	return resp.Send()
}

// NoContent envía una respuesta con código 204 (sin cuerpo)
func NoContent(c echo.Context, message string) error {
	if message == "" {
		message = "No Content"
	}
	return New(c).Message(message).Code(http.StatusNoContent).Send()
}

// UserError envía una respuesta con código 400
func UserError(c echo.Context, message string, data interface{}) error {
	if message == "" {
		message = "Bad Request"
	}
	// Si no hay datos, se deja el valor por defecto (vacío)
	resp := New(c).Message(message).Code(http.StatusBadRequest)
	if data != nil {
		resp.Data(data)
	}
	return resp.Send()
}

// Unauthorized envía una respuesta con código 401
func Unauthorized(c echo.Context, message string, data interface{}) error {
	if message == "" {
		message = "Unauthorized"
	}
	// Si no hay datos, se deja el valor por defecto (vacío)
	resp := New(c).Message(message).Code(http.StatusUnauthorized)
	if data != nil {
		resp.Data(data)
	}
	return resp.Send()
}

// Forbidden envía una respuesta con código 403
func Forbidden(c echo.Context, message string, data interface{}) error {
	if message == "" {
		message = "Forbidden"
	}
	// Si no hay datos, se deja el valor por defecto (vacío)
	resp := New(c).Message(message).Code(http.StatusForbidden)
	if data != nil {
		resp.Data(data)
	}
	return resp.Send()
}

// NotFound envía una respuesta con código 404
func NotFound(c echo.Context, message string, data interface{}) error {
	if message == "" {
		message = "Not Found"
	}
	// Si no hay datos, se deja el valor por defecto (vacío)
	resp := New(c).Message(message).Code(http.StatusNotFound)
	if data != nil {
		resp.Data(data)
	}
	return resp.Send()
}

// NotFound envía una respuesta con código 422
func Unprocessable(c echo.Context, message string, data interface{}) error {
	if message == "" {
		message = "Unprocessable Content"
	}
	// Si no hay datos, se deja el valor por defecto (vacío)
	resp := New(c).Message(message).Code(http.StatusUnprocessableEntity)
	if data != nil {
		resp.Data(data)
	}
	return resp.Send()
}

// ServerError envía una respuesta con código 500
func ServerError(c echo.Context, message string, data interface{}) error {
	if message == "" {
		message = "Internal Server Error"
	}
	// Si no hay datos, se deja el valor por defecto (vacío)
	resp := New(c).Message(message).Code(http.StatusInternalServerError)
	if data != nil {
		resp.Data(data)
	}
	return resp.Send()
}
