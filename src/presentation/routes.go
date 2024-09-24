package presentation

import (
	"FileManager/src/presentation/auth"
	"FileManager/src/presentation/files"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas las rutas
func SetupRoutes(router *gin.Engine) {

	// Crear un grupo para las rutas API
	apiGroup := router.Group("/api/v1")

	// Configurar las rutas de autenticación
	auth.SetupAuthRoutes(apiGroup) // Pasar el grupo a las rutas de auth

	// Configurar las rutas de archivos
	files.SetupFilesRoutes(apiGroup) // Pasar el grupo a las rutas de files

	// Otras rutas adicionales pueden añadirse aquí
}
