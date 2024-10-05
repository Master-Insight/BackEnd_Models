package middlewares

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// LoggerMiddleware función intermediaria para retornar el middleware Logger del framework Echo
func LoggerMiddleware() echo.MiddlewareFunc {
	myLog, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("No se pudo abrir o crear el archivo de logs")
	}

	logConfig := middleware.LoggerConfig{
		Output: myLog,
		Format: "ip=${remote_ip}, time=${time_rfc3339}, method=${method}, status=${status}, uri=${uri}, error=${error}, \n",
	}
	//defer myLog.Close()

	return middleware.LoggerWithConfig(logConfig)
}
