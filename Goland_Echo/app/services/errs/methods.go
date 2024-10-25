package errs

import "net/http"

// Definimos una función genérica para crear errores de cliente (4xx)
func NewClientError(message string, description string, codeHTTP int) *Error {
	return New(message, description, codeHTTP)
}

// Definimos una función genérica para crear errores del servidor (5xx)
func NewServerError(message string, description string, codeHTTP int) *Error {
	return New(message, description, codeHTTP)
}

// * Errores comunes del cliente (4xx) -----------------------------------------------------------------------------------------------

// Errores comunes del cliente (4xx)
func NewBadRequest(message, description string) *Error {
	return NewClientError(message, description, http.StatusBadRequest)
}

func NewUnauthorized(message, description string) *Error {
	return NewClientError(message, description, http.StatusUnauthorized)
}

// Puedes seguir este patrón para todos los demás errores

// * Errores comunes del servidor (5xx) ----------------------------------------------------------------------------------------------

// Errores comunes del servidor (5xx)
func NewInternalServerError(message, description string) *Error {
	return NewServerError(message, description, http.StatusInternalServerError)
}
