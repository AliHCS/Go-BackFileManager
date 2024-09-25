package presentation

import (
	"FileManager/src/presentation/auth"
	"FileManager/src/presentation/files"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRoutes configura todas las rutas
func SetupRoutes(router *gin.Engine, client *mongo.Client) {

	// Crear un grupo para las rutas API
	apiGroup := router.Group("/api/v1")

	// Configurar las rutas de autenticación
	auth.SetupAuthRoutes(apiGroup, client) // Pasar el grupo y el cliente a las rutas de auth

	// Configurar las rutas de archivos
	files.SetupFilesRoutes(apiGroup) // Asegúrate de que esto también esté configurado correctamente

	// Otras rutas adicionales pueden añadirse aquí
}
