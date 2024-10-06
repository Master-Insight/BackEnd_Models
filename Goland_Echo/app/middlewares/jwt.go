package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// JWTMiddleware función intermediaria para retornar el middleware JWT del framework Echo
func JWTMiddleware() echo.MiddlewareFunc {

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{

		// * Secret Key
		SigningKey: []byte("secret"), // ! auth.GetPublicKey(),

		// * Signing method
		// Optional. Default value HS256.
		SigningMethod: "RS256",

		// *Context key
		// to store user information from the token into context.
		// Optional. Default value "user".
		ContextKey: "auth_token",

		// *TokenLookup
		// is a string in the form of "<source>:<name>" or "<source>:<name>,<source>:<name>" that is used
		// to where extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>" or "header:<name>:<cut-prefix>"
		// 			`<cut-prefix>` is argument value to cut/trim prefix of the extracted value. This is useful if header
		//			value has static prefix like `Authorization: <auth-scheme> <authorisation-parameters>` where part that we
		//			want to cut is `<auth-scheme> ` note the space at the end.
		//			In case of JWT tokens `Authorization: Bearer <token>` prefix we cut is `Bearer `.
		// If prefix is left empty the whole value is returned.
		// - "query:<name>"
		// - "param:<name>"
		// - "cookie:<name>"
		// - "form:<name>"
		// Multiple sources example:
		// - "header:Authorization:Bearer ,cookie:myowncookie"
		// ? TokenLookup: "header:Authorization",

		// * AuthScheme
		// Optional. Default value "Bearer".
		// AuthScheme: "Bearer",

		// * Claims
		// are extendable claims data defining token content. Used by default ParseTokenFunc implementation.
		// Not used if custom ParseTokenFunc is set.
		// Optional. Defaults to function returning jwt.MapClaims
		// ! Claims:        &auth.JWTCustomClaims{},
	}

	return middleware.JWTWithConfig(config)
}
