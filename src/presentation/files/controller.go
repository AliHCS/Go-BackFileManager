package files

import (
	"FileManager/src/presentation/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FileController maneja las solicitudes de archivos
type FileController struct {
	fileService *services.FileService
}

// Upload maneja la solicitud de carga de archivos
func (f *FileController) Upload(c *gin.Context) {
	// Lógica de carga de archivos
	userId := c.PostForm("userID")
	fmt.Println(userId)
	// Obtenemos el archivo desde el formulario

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Error al obtener el archivo: %s", err.Error())
		return
	}
	// Abrir el archivo
	fileContent, err := file.Open()
	if err != nil {
		c.String(http.StatusBadRequest, "Error al abrir el archivo: %s", err.Error())
		return
	}
	defer fileContent.Close()

	// Llamar al servicio
	filePath, err := f.fileService.UploadFile(fileContent, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error al guardar el archivo: %s", err.Error())
		return
	}
	// Devolver respuesta
	c.JSON(http.StatusOK, gin.H{
		"message":  "Archivo cargado exitosamente",
		"filePath": filePath,
	})
}

// List maneja la solicitud de lista de archivos
func (f *FileController) List(c *gin.Context) {
	// Lógica para listar archivos
	c.String(200, "Lista de archivos")
}
