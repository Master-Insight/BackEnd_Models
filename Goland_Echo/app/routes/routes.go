package routes

import "github.com/labstack/echo/v4"

// Routes carga las rutas de la aplicación
func Routes(e *echo.Echo) {
	// Agregar las rutas de cada módulo
	e.GET("/api/example", ExampleHandler)
}

// Ejemplo de controlador
func ExampleHandler(c echo.Context) error {
	return c.JSON(200, map[string]string{"message": "This is an example route"})
}
