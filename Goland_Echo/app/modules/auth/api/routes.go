package auth_api

import (
	"github.com/labstack/echo/v4"
)

// AuthRoutes configura las rutas de autenticación
func AuthRoutes(e *echo.Group) {
	group := e.Group("/auth")

	group.POST("/login", loginHandler)
	group.POST("/register", registerHandler)
}
