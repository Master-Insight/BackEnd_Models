package env

import "log"

// LoadEnviroment decide qué método usar para cargar la configuración
func LoadEnviroment(method string) *ConfigEnv {
	switch method {
	case "env":
		return LoadEnv()
	case "json":
		return LoadJsonEnv()
	default:
		log.Fatalf("Método de carga de configuración no soportado: %v", method)
		return nil
	}
}
