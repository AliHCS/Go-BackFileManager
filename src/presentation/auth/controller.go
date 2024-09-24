package auth

import "github.com/gin-gonic/gin"

// AuthController maneja las solicitudes de autenticación
type AuthController struct{}

// Login maneja la solicitud de inicio de sesión
func (a *AuthController) Login(c *gin.Context) {
	// Lógica de inicio de sesión
	c.String(200, "Iniciar sesión exitoso")
}

// Register maneja la solicitud de registro
func (a *AuthController) Register(c *gin.Context) {
	// Lógica de registro
	c.String(200, "Registro exitoso")
}
