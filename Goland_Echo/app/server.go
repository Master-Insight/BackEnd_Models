package app

import (
	"aprendiendoGo/app/middlewares"
	"aprendiendoGo/app/routes"
	"aprendiendoGo/app/services/env"
	"log"

	"github.com/labstack/echo/v4"
)

// Boot inicia la api
func Boot() {

	// App initialization ------------------------------
	e := echo.New()

	// App Configurations --------------------------------
	cfg := env.LoadEnviroment("json")
	port := cfg.Config.Port

	// App Middleware --------------------------------
	e.Use(middlewares.LoggerMiddleware())
	e.Use(middlewares.RecoverMiddleware())
	e.Use(middlewares.CORSMiddleware(cfg.CorsOrigin[0]))

	// App Routes --------------------------------
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Hello, Go Echo Backend!"})
	})
	// Inicializar rutas de la aplicación
	routes.Routes(e)

	// App Launch --------------------------------
	log.Println("Server running on port:", port)
	e.Logger.Fatal(e.Start(":" + port))
}
