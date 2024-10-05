package env

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Carga de archivo Env
func LoadEnv() *ConfigEnv {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Asignar valores de entorno a la configuración
	return &ConfigEnv{
		// la alternativa es os.Getenv("PORT") pero no te asegura que este
		Port:                 getEnv("PORT", "8080"),
		CorsOrigin:           getEnv("CORS_ORIGIN", "*"),
		MongoURI:             getEnv("MONGO_URI", ""),
		JWTSecretCode:        getEnv("JWT_SECRET_CODE", ""),
		AdminUsers:           strings.Split(getEnv("USERS_ADMIN", ""), ","),
		AdminPassword:        getEnv("USER_ADMIN_PASS", ""),
		GmailUserName:        getEnv("GMAIL_USER_NAME", ""),
		GmailUserApp:         getEnv("GMAIL_USER_APP", ""),
		GmailPassApp:         getEnv("GMAIL_PASS_APP", ""),
		LinkedInClientID:     getEnv("LINKEDIN_CLIENT_ID", ""),
		LinkedInClientSecret: getEnv("LINKEDIN_CLIENT_SECRET", ""),
		LinkedInRedirectURI:  getEnv("LINKEDIN_REDIRECT_URI", ""),
		LinkedInScope:        getEnv("LINKEDIN_SCOPE", ""),
		CloudinaryName:       getEnv("CLOUDINARY_NAME", ""),
		CloudinaryKey:        getEnv("CLOUDINARY_KEY", ""),
		CloudinarySecret:     getEnv("CLOUDINARY_SECRET", ""),
		CloudinaryURL:        getEnv("CLOUDINARY_URL", ""),
	}
}

// getEnv obtiene el valor de una variable de entorno o un valor por defecto si no existe
func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
