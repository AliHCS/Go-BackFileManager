package files

import (
	"github.com/gin-gonic/gin"
)

// SetupFilesRoutes define las rutas para archivos
func SetupFilesRoutes(router *gin.RouterGroup) {

	// Crear una instancia del controlador de archivos
	fileController := &FileController{}
	// Definir las rutas bajo el grupo
	router.POST("/files/upload", fileController.Upload)
	router.POST("/files/getAll", fileController.List)

}
