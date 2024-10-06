package env

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Carga de archivo Env
func LoadEnv() *ConfigEnvStruc {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Convertir las variables de entorno a un slice de strings
	corsOrigin := strings.Split(getEnv("CORS_ORIGIN", "*"), ",")

	// Asignar valores de entorno a la configuración
	return &ConfigEnvStruc{
		// la alternativa es os.Getenv("PORT") pero no te asegura que este
		Config: Config{
			AppName: getEnv("APP_NAME", "Go Echo Backend"),
			Port:    getEnv("PORT", "8080"),
		},
		CorsOrigin: corsOrigin,
		Codes: Codes{
			JWT: getEnv("JWT_SECRET_CODE", ""),
		},
		Admin: Admin{
			Users: []string{getEnv("USERS_ADMIN", "Admin")},
			Pass:  getEnv("USER_ADMIN_PASS", ""),
		},
		Services: Services{
			Persistence: Persistence{
				Service: getEnv("PERSISTENCE", ""),
				Mongo: []Mongo{
					{URI: getEnv("MONGO_URI", "")},
				},
			},
			Email: Email{
				UserName:  getEnv("USER_NAME", ""),
				GmailUser: getEnv("GMAIL_USER", ""),
				GmailPass: getEnv("GMAIL_PASS", ""),
			},
			Images: Images{
				Cloudinary: Cloudinary{
					Name:          getEnv("CLOUDINARY_NAME", ""),
					Key:           getEnv("CLOUDINARY_KEY", ""),
					Secret:        getEnv("CLOUDINARY_SECRET", ""),
					URL:           getEnv("CLOUDINARY_URL", ""),
					DefaultFolder: getEnv("CLOUDINARY_DEFAULT_FOLDER", ""),
				},
			},
		},
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
