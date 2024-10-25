package middlewares

import (
	"aprendiendoGo/app/services/errs"
	"aprendiendoGo/app/services/response"

	"github.com/labstack/echo/v4"
)

// ErrorHandlerMiddleware captura y maneja los errores
func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			if customErr, ok := err.(*errs.Error); ok {
				return response.GenericResponse(c, customErr.CodeHTTP, nil, customErr.Message)
			}
			return response.ServerError(c, "Internal server error", nil)
		}
		return nil
	}
}
