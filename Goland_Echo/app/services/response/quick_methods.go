package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ResponseTypes mapea códigos de estado con mensajes predeterminados
var ResponseTypes = map[int]string{
	http.StatusOK:                  "Success",
	http.StatusCreated:             "Created",
	http.StatusNoContent:           "No Content",
	http.StatusBadRequest:          "Bad Request",
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusForbidden:           "Forbidden",
	http.StatusNotFound:            "Not Found",
	http.StatusUnprocessableEntity: "Unprocessable Content",
	http.StatusInternalServerError: "Internal Server Error",
}

// GenericResponse genera y envía una respuesta genérica
func GenericResponse(c echo.Context, code int, data interface{}, message string) error {
	if message == "" {
		message = ResponseTypes[code]
	}
	resp := New(c).Message(message).Code(code)
	if data != nil {
		resp.Data(data)
	}
	return resp.Send()
}

// * Ejemplo de uso:
// func handlerSuccess(c echo.Context) error {
// 	userData := map[string]string{
// 			"username": "JohnDoe",
// 			"email": "john@example.com",
// 	}
// 	return response.Success(c, userData, "Operación exitosa")
// }

// Success envía una respuesta 200 OK
func Success(c echo.Context, data interface{}, message string) error {
	return GenericResponse(c, http.StatusOK, data, message)
}

// Created envía una respuesta 201 Created
func Created(c echo.Context, data interface{}, message string) error {
	return GenericResponse(c, http.StatusCreated, data, message)
}

// NoContent envía una respuesta 204 No Content
func NoContent(c echo.Context, message string) error {
	return GenericResponse(c, http.StatusNoContent, nil, message)
}

// UserError envía una respuesta 400 Bad Request
func UserError(c echo.Context, message string, data interface{}) error {
	return GenericResponse(c, http.StatusBadRequest, data, message)
}

// Unauthorized envía una respuesta 401 Unauthorized
func Unauthorized(c echo.Context, message string, data interface{}) error {
	return GenericResponse(c, http.StatusUnauthorized, data, message)
}

// Forbidden envía una respuesta 403 Forbidden
func Forbidden(c echo.Context, message string, data interface{}) error {
	return GenericResponse(c, http.StatusForbidden, data, message)
}

// NotFound envía una respuesta 404 Not Found
func NotFound(c echo.Context, message string, data interface{}) error {
	return GenericResponse(c, http.StatusNotFound, data, message)
}

// Unprocessable envía una respuesta 422 Unprocessable Entity
func Unprocessable(c echo.Context, message string, data interface{}) error {
	return GenericResponse(c, http.StatusUnprocessableEntity, data, message)
}

// ServerError envía una respuesta 500 Internal Server Error
func ServerError(c echo.Context, message string, data interface{}) error {
	return GenericResponse(c, http.StatusInternalServerError, data, message)
}
