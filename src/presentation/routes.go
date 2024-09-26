package presentation

import (
	"FileManager/src/presentation/auth"
	"FileManager/src/presentation/files"
	"FileManager/src/presentation/middlewares" // Importa el paquete de middlewares

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRoutes configura todas las rutas
func SetupRoutes(router *gin.Engine, client *mongo.Client) {

	// Crear un grupo para las rutas API
	apiGroup := router.Group("/api/v1")

	// Aplicar el middleware ParseBodyMiddleware a todas las rutas en apiGroup
	apiGroup.Use(middlewares.ParseBodyMiddleware())

	// Configurar las rutas de autenticación
	auth.SetupAuthRoutes(apiGroup, client) // Pasar el grupo y el cliente a las rutas de auth

	// Configurar las rutas de archivos
	files.SetupFilesRoutes(apiGroup, client) // Asegúrate de que esto también esté configurado correctamente

	// Otras rutas adicionales pueden añadirse aquí
}
