package env

import (
	"log"
	"sync"
)

var once sync.Once
var configEnv *ConfigEnvStruc

// Get devuelve la configuración cargada
func Get() *ConfigEnvStruc {
	if configEnv == nil {
		log.Fatalf("La configuración no ha sido cargada. Debes invocar LoadEnvironment primero.")
	}
	return configEnv
}

// LoadEnvironment decide qué método usar para cargar la configuración
func LoadEnvironment(method string) {
	once.Do(func() {
		switch method {
		case "env":
			configEnv = loadEnvFile()
		case "json":
			configEnv = loadJsonFile()
		default:
			log.Fatalf("Método de carga de configuración no soportado: %v", method)
		}
	})
} // Ejemplo: env.LoadEnvironment("json")
