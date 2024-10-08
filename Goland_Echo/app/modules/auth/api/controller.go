package auth_api

import (
	auth_data "aprendiendoGo/app/modules/auth/data"
	auth_logic "aprendiendoGo/app/modules/auth/logic"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Estructura para el controlador
var validate = validator.New()

// Controlador para el login

func loginHandler(c echo.Context) error {
	// Aquí puedes agregar la lógica para manejar el login
	// Por ahora, solo devolveremos un mensaje simple.
	return c.JSON(http.StatusOK, map[string]string{"message": "Esto es Login Auth"})
}

func registerHandler(c echo.Context) error {
	var user auth_data.User

	// Bind JSON a la estructura de usuario
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error al analizar los datos"})
	}

	// Validar los datos
	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Validación fallida", "errors": err.Error()})
	}

	// Llamar al servicio de registro y obtener la respuesta
	registeredUser := auth_logic.RegisterUser(user)

	// Devolver el usuario registrado
	return c.JSON(http.StatusOK, registeredUser)
}
