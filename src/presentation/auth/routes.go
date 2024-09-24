package auth

import (
	"FileManager/src/presentation/services"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes define las rutas de autenticación
func SetupAuthRoutes(router *gin.RouterGroup) {
	// Crear una instancia del servicio de autenticación
	authService := services.NewAuthService()

	// Crear una instancia del controlador de autenticación
	authController := &AuthController{authService}
	// Definir las rutas bajo el grupo
	router.POST("/auth/login", authController.Login)
	router.POST("/auth/register", authController.Register)

	// Puedes añadir más rutas aquí, por ejemplo:
	// router.POST("/login", loginHandler)
}
