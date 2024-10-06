package routes

import "github.com/labstack/echo/v4"

// StaticRoutes crea las rutas estáticas de la aplicación
func StaticRoutes(e *echo.Echo) {
	e.Static("/", "public")
	// se usa de manera similar "e.File" cuando se quiere mostrar 1 solo archivo
}
