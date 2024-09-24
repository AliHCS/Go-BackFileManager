package presentation

import (
	"FileManager/src/config"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	// Cargar las variables de entorno desde el archivo .env

	cfg := config.LoadEnv()

	// Crear una nueva instancia de Gin
	router := gin.Default()

	// Configurar las rutas
	SetupRoutes(router)

	// Iniciar el servidor
	log.Printf("Servidor corriendo en el puerto %s", cfg.PORT)
	if err := router.Run(":" + cfg.PORT); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
