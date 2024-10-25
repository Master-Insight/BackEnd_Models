package auth_api

import (
	auth_data "aprendiendoGo/app/modules/auth/data"
	auth_logic "aprendiendoGo/app/modules/auth/logic"
	"aprendiendoGo/app/services/response"
	"aprendiendoGo/app/services/validatorpartner"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Estructura para el controlador
//var validate = validator.New()

// Controller - Login
func loginHandler(c echo.Context) error {
	// Aquí puedes agregar la lógica para manejar el login
	// Por ahora, solo devolveremos un mensaje simple.
	return c.JSON(http.StatusOK, map[string]string{"message": "Esto es Login Auth"})
}

// Controller - Register
func registerHandler(c echo.Context) error {
	var user auth_data.User

	// Bind JSON a la estructura de usuario
	if err := c.Bind(&user); err != nil {
		return response.UserError(c, "Error al analizar los datos", nil)
	}

	// * Validar los datos
	// Crear una nueva instancia de validador
	validator := validatorpartner.New().Data(user)

	// Validar los datos y gestionar respuesta en caso de errores
	if validationErrors := validator.ValidateSend(c); validationErrors != nil {
		// Si hay errores de validación, se retornan
		return validationErrors
	}

	// Llamar al servicio de registro y obtener la respuesta
	registeredUser, err := auth_logic.RegisterUser(user)
	if err != nil {
		return response.UserError(c, "Error al registrar el usuario", err)
	}

	// Devolver el usuario registrado
	return response.Created(c, registeredUser, "User created")
}
