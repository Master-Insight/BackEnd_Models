package env

import "log"

var configEnv *ConfigEnvStruc

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
