package auth

import (
	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes define las rutas de autenticación
func SetupAuthRoutes(router *gin.RouterGroup) {

	// Crear una instancia del controlador de autenticación
	authController := &AuthController{}
	// Definir las rutas bajo el grupo
	router.POST("/auth/login", authController.Login)
	router.POST("/auth/register", authController.Register)

	// Puedes añadir más rutas aquí, por ejemplo:
	// router.POST("/login", loginHandler)
}
