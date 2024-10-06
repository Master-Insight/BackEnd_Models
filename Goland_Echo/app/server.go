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

	// App initialization ------------------------------
	e := echo.New()

	// App Configurations --------------------------------
	cfg := env.LoadEnviroment("json")
	port := cfg.Config.Port

	// App Data Source Configuration --------------------------------
	mongo.GetMongoInstance(cfg.Services.Persistence.Mongo[0].URI)
	defer mongo.DisconnectMongo()

	// App Middleware --------------------------------
	e.Use(middlewares.LoggerMiddleware())
	e.Use(middlewares.RecoverMiddleware())
	e.Use(middlewares.CORSMiddleware(cfg.CorsOrigin[0]))

	// App Routes --------------------------------
	modules.Routes(e)

	// App Launch --------------------------------
	log.Println("Server running on port:", port)
	e.Logger.Fatal(e.Start(":" + port))
}
