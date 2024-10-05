package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORSMiddleware función intermediaria para retornar el middleware CORS del framework Echo
func CORSMiddleware(corsOrigin string) echo.MiddlewareFunc {

	return middleware.CORSWithConfig(middleware.CORSConfig{
		// Skipper defines a function to skip middleware.
		Skipper: middleware.DefaultSkipper,

		AllowOrigins: []string{corsOrigin}, // default {"*"}

		// Optional. Default value DefaultCORSConfig.AllowMethods.
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		// alternative --> AllowMethods: []string{"GET", "HEAD", "PUT", "POST", "DELETE", "OPTIONS"},

		// AllowHeaders defines a list of request headers that can be used
		AllowHeaders: []string{}, // default

		// AllowCredentials indicates whether or not the response to the request
		// can be exposed when the credentials flag is true. When used as part of
		// a response to a preflight request, this indicates whether or not the
		// actual request can be made using credentials.
		// Optional. Default value false.
		AllowCredentials: false,

		// ExposeHeaders defines a whitelist headers
		ExposeHeaders: []string{}, // default

		// MaxAge indicates how long (in seconds) the results can be cached.
		MaxAge: 0, // default
	})
}
