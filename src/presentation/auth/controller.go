package auth

import (
	"FileManager/src/domain/dtos/auth"
	"FileManager/src/presentation/services"
	"fmt"
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
	// Obtener los props del contexto
	props, exists := c.Get("body")
	// Imprimir props para depuración
	fmt.Println("Props recibidos:", props)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cuerpo de la solicitud no encontrado",
		})
		return
	}
	// Intentar crear un nuevo LoginDTO usando los datos recibidos
	dto, err := auth.NewLoginDTO(props.(map[string]string))
	if err != nil {
		// Si hay algún error en la creación del DTO, devolver una respuesta con el error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Llamar al servicio de registro de usuario
	message, err := a.authService.LoginUser(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

// Register maneja la solicitud de registro
func (a *AuthController) Register(c *gin.Context) {
	// Lógica de registro
	// Obtener los props del contexto
	props, exists := c.Get("body")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cuerpo de la solicitud no encontrado",
		})
		return
	}

	// Intentar crear un nuevo RegisterDTO usando los datos recibidos
	dto, err := auth.NewRegisterDTO(props.(map[string]string))

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
		"response": message,
	})
}
