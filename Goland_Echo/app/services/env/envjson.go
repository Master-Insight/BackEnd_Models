package env

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var once sync.Once      // Lo utilizamos para construir un singleton
var env *ConfigEnvStruc // Variable Una sola instancia

// Carga de archivo Env
func LoadJsonEnv() *ConfigEnvStruc {

	once.Do(func() { // Función nativa que hace que el código se ejecute una sola vez

		file, err := os.Open("env.json")
		if err != nil {
			log.Fatalf("Error al abrir el archivo env.json: %v", err)
		}
		defer file.Close()

		// Leer el archivo
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatalf("Error al leer el archivo env.json: %v", err)
		}

		// Convertir JSON a la estructura Env
		err = json.Unmarshal(bytes, &env)
		if err != nil {
			log.Fatalf("Error al deserializar env.json: %v", err)
		}

		log.Println("Archivo de configuración cargado correctamente")

	})

	return env
}
