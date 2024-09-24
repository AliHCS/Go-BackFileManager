package auth

import (
	"FileManager/src/domain/dtos/auth"
	"FileManager/src/presentation/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController maneja las solicitudes de autenticación
type AuthController struct {
	authService *services.AuthService // Agregar AuthService como un campo
}

// Login maneja la solicitud de inicio de sesión
func (a *AuthController) Login(c *gin.Context) {
	// Lógica de inicio de sesión
	c.String(200, "Iniciar sesión exitoso")
}

// Register maneja la solicitud de registro
func (a *AuthController) Register(c *gin.Context) {
	// Lógica de registro
	// Crear un mapa para almacenar los datos del request body
	var props map[string]string

	// Parsear el cuerpo de la solicitud en formato JSON al mapa
	if err := c.BindJSON(&props); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos en el cuerpo de la solicitud",
		})
		return
	}

	// Intentar crear un nuevo RegisterDTO usando los datos recibidos
	dto, err := auth.NewRegisterDTO(props)
	if err != nil {
		// Si hay algún error en la creación del DTO, devolver una respuesta con el error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Llamar al servicio de registro de usuario
	message, err := a.authService.RegisterUser(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"user": gin.H{
			"email": dto.Email,
			"name":  dto.Name,
		},
	})
}
