package app

import (
	"aprendiendoGo/app/middlewares"
	"aprendiendoGo/app/modules"
	"aprendiendoGo/app/services/db/mongo"
	"aprendiendoGo/app/services/env"
	"log"

	"github.com/labstack/echo/v4"
)

// Boot inicia la api
func Boot() {

	// App initialization --------------------------------
	e := echo.New()

	// App Configurations --------------------------------
	env.LoadEnvironment("json") // Carga el ambiente una sola vez
	cfg := env.Get()
	port := cfg.Config.Port

	// Data Source Config --------------------------------
	mongo.GetMongoInstance(cfg.Services.Persistence.Mongo[0].URI)
	defer mongo.DisconnectMongo()

	// Middleware ----------------------------------------
	e.Use(middlewares.LoggerMiddleware())
	e.Use(middlewares.RecoverMiddleware())
	e.Use(middlewares.CORSMiddleware(cfg.CorsOrigin[0]))
	//e.Use(middlewares.ErrorHandlerMiddleware)

	// Routes --------------------------------------------
	modules.Routes(e)

	// App Launch ----------------------------------------
	log.Printf("Server running on port: %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
