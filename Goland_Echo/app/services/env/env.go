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

// LoadEnviroment decide qué método usar para cargar la configuración
func LoadEnviroment(method string) *ConfigEnvStruc {
	switch method {
	case "env":
		configEnv = LoadEnv()
	case "json":
		configEnv = LoadJsonEnv()
	default:
		log.Fatalf("Método de carga de configuración no soportado: %v", method)
		return nil
	}

	return configEnv
}
