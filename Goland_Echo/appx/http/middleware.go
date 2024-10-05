package http

import (
	"aprendiendoGo/app/http/middlewares"

	"github.com/labstack/echo/v4"
)

// Middlewares carga todos los middlewares de la aplicación
func GenericMiddleware(e *echo.Echo) {
	e.Use(middlewares.LoggerMiddleware())
	e.Use(middlewares.RecoverMiddleware())
	e.Use(middlewares.CORSMiddleware())
}
