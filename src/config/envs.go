package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Configura la estructura para almacenar las variables de entorno
type Config struct {
	PORT string
}

// CargarEnv carga las variables de entorno desde el archivo .env
func LoadEnv() *Config {
	// Cargar las variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Crear una nueva instancia de Config y cargar las variables
	return &Config{
		PORT: getEnv("PORT", "8080"),
	}
}

// getEnv obtiene el valor de la variable de entorno o un valor por defecto
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
