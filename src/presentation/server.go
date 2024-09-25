package presentation

import (
	"FileManager/src/config"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	// Cargar las variables de entorno desde el archivo .env
	cfg := config.LoadEnv()

	// Crear una nueva instancia de Gin
	router := gin.Default()

	// Construir el URI de MongoDB
	uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017/", cfg.MONGO_USERNAME, cfg.MONGO_PASSWORD)

	client, err := config.ConnectToMongoDB(uri)
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil { // Usar un contexto válido aquí
			log.Fatalf("Error al desconectar de MongoDB: %v", err)
		}
	}()
	log.Println("Conexión exitosa a MongoDB")

	// Configurar las rutas
	SetupRoutes(router, client) // Aquí pasamos el cliente de MongoDB

	// Iniciar el servidor
	gin.SetMode(cfg.GIN_MODE)

	log.Printf("Servidor corriendo en el puerto %s", cfg.PORT)
	if err := router.Run(":" + cfg.PORT); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
