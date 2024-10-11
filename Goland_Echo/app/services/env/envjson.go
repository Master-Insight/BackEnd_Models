package env

import (
	"encoding/json"
	"log"
	"os"
)

func loadJsonFile() *ConfigEnvStruc {
	var config ConfigEnvStruc
	file, err := os.Open("env.json")
	if err != nil {
		log.Fatalf("Error al abrir el archivo env.json: %v", err)
	}
	defer file.Close()

	// Leer el archivo
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("Error al deserializar env.json: %v", err)
	}

	log.Println("Archivo de configuración cargado correctamente")
	return &config
}
