package env

// Estructura de las variables de entorno de la aplicación partiendo desde un Json
type ConfigEnvStruc struct {
	Config     Config   `json:"config"`
	CorsOrigin []string `json:"cors_origin"`
	Codes      Codes    `json:"codes"`
	Admin      Admin    `json:"admin"`
	Services   Services `json:"services"`
}

// AppConfig contiene la configuración de la aplicación
type Config struct {
	AppName string `json:"name"`
	Port    string `json:"port"`
}

// Codes contiene los códigos sensibles como JWT
type Codes struct {
	JWT string `json:"jwt"`
}

// Admin contiene la información de los administradores
type Admin struct {
	Users []string `json:"users"`
	Pass  string   `json:"pass"`
}

// Services contiene las configuraciones de los servicios
type Services struct {
	Persistence Persistence `json:"persistence"`
	Email       Email       `json:"email"`
	Images      Images      `json:"images"`
}

// * SERVICES:
// --------------------------------

// Persistence contiene la configuración de persistencia
type Persistence struct {
	Service string  `json:"service"`
	Mongo   []Mongo `json:"mongo"`
}

// Persistence --> Mongo estructura para la conexión a MongoDB
type Mongo struct {
	URI      string `json:"uri"`
	Database string `json:"database"`
}

// Email contiene la configuración del servicio de email
type Email struct {
	UserName  string `json:"user_name"`
	GmailUser string `json:"gmail_user"`
	GmailPass string `json:"gmail_pass"`
}

// Image contiene la configuración del servicio de imágenes
type Images struct {
	Cloudinary Cloudinary `json:"cloudinary"`
}

// Cloudinary contiene la configuración de Cloudinary
type Cloudinary struct {
	Name          string `json:"name"`
	Key           string `json:"key"`
	Secret        string `json:"secret"`
	URL           string `json:"url"`
	DefaultFolder string `json:"default_folder"`
}
