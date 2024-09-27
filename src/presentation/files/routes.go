package files

import (
	"FileManager/src/presentation/middlewares"
	"FileManager/src/presentation/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo" // Importa mongo para pasar el cliente
)

// SetupFilesRoutes define las rutas para archivos
func SetupFilesRoutes(router *gin.RouterGroup, client *mongo.Client) {
	// Crear una instacia del servicio
	fileService := services.NewFilesService(client)

	// Crear una instancia del controlador de archivos
	fileController := &FileController{fileService}

	// Aplicar el middleware AuthMiddleware a las rutas de archivos
	authMiddleware := middlewares.AuthMiddleware(client) // Crear una instancia del middleware con el cliente de MongoDB

	// Definir las rutas bajo el grupo y aplicar el middleware
	router.POST("/files/upload", authMiddleware, fileController.Upload) // Rutas protegidas por el middleware
	router.POST("/files/getAll", authMiddleware, fileController.List)   // Rutas protegidas por el middleware
}
