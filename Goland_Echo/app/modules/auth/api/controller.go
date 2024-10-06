package auth_api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func loginHandler(c echo.Context) error {
	// Aquí puedes agregar la lógica para manejar el login
	// Por ahora, solo devolveremos un mensaje simple.
	return c.JSON(http.StatusOK, map[string]string{"message": "Esto es Login Auth"})
}

func registerHandler(c echo.Context) error {
	// Aquí puedes agregar la lógica para manejar el login
	// Por ahora, solo devolveremos un mensaje simple.
	return c.JSON(http.StatusOK, map[string]string{"message": "Esto es Register Auth"})
}
