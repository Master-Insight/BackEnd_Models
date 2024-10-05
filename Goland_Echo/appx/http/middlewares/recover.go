package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RecoverMiddleware función intermediaria para retornar el middleware Recover del framework Echo
func RecoverMiddleware() echo.MiddlewareFunc {

	return middleware.Recover()

}
