package files

import (
	"FileManager/src/domain/dtos/files"
	"FileManager/src/presentation/services"
	"fmt"
	"net/http"
	"time"

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

	// Crear una instancia de FileInfo
	fileInfo := files.FileInfo{
		OriginalName: file.Filename,
		Filename:     fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename), // Generar un nombre único
		Path:         fmt.Sprintf("uploads/%s", file.Filename),               // Definir la ruta de almacenamiento
		MimeType:     file.Header.Get("Content-Type"),
		Size:         file.Size,
	}

	dto, err := files.NewUploadFileDto(userId, fileInfo)
	if err != nil {
		c.String(http.StatusBadRequest, "Error al crear el DTO: %s", err.Error())
		return
	}

	// Llamar al servicio
	response, err := f.fileService.UploadFile(fileContent, dto)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error al guardar el archivo: %s", err.Error())
		return
	}
	// Devolver respuesta
	c.JSON(http.StatusOK, gin.H{
		"message":  "Archivo cargado exitosamente",
		"response": response,
	})
}

// List maneja la solicitud de lista de archivos
func (f *FileController) List(c *gin.Context) {
	// Llamar al servicio
	response, err := f.fileService.GetAllFiles()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error al guardar el archivo: %s", err.Error())
		return
	}
	// Devolver respuesta
	c.JSON(http.StatusOK, gin.H{
		"message":  "Consulta Exitosa",
		"response": response,
	})
}
