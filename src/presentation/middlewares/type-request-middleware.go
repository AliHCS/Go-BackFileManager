package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ParseBodyMiddleware es un middleware que intenta parsear el cuerpo de la solicitud a un mapa
func ParseBodyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		props := make(map[string]string)

		// Verificar si el tipo de contenido es multipart/form-data
		if c.ContentType() == "multipart/form-data" {
			// Obtener los datos del formulario
			if err := c.ShouldBind(&props); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Datos inválidos en el cuerpo de la solicitud",
				})
				c.Abort() // Abortamos la cadena de middleware
				return
			}
		} else {
			// Intentar leer el cuerpo en formato JSON
			if err := c.ShouldBindJSON(&props); err != nil {
				// Si falla el JSON, intentar leer en formato x-www-form-urlencoded
				if err := c.ShouldBind(&props); err != nil {
					// Si falla también el formato x-www-form-urlencoded, devolver error
					c.JSON(http.StatusBadRequest, gin.H{
						"error": "Datos inválidos en el cuerpo de la solicitud",
					})
					c.Abort() // Abortamos la cadena de middleware
					return
				}
			}
		}

		// Almacenar el mapa en el contexto para usarlo más tarde
		c.Set("body", props)
		// Imprimir los props para depuración
		fmt.Println("Props parseados:", props)

		// Continuar al siguiente middleware o manejador
		c.Next()
	}
}
