package modules

import (
	auth_api "aprendiendoGo/app/modules/auth/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Routes carga las rutas de la aplicación
func Routes(e *echo.Echo) {

	// Ruta que muestra que funciona
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, Go Echo Backend!"})
	})

	// Rutas de la API V1
	v1 := e.Group("/v1")
	auth_api.AuthRoutes(v1)

	// Manejo de error 404
	e.Any("*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Error 404 -This route is not allowed"})
	})
}
