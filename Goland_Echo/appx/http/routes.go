package http

import (
	"aprendiendoGo/app/http/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler / Controller
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "hello"})
}

// Routes carga las rutas de la aplicación
func Routes(e *echo.Echo) {
	e.GET("/", hello)

	e.GET("/v1/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, map[string]string{
			"id": id,
		})
	})

	routes.StaticRoutes(e)
}
