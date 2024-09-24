package files

import "github.com/gin-gonic/gin"

// FileController maneja las solicitudes de archivos
type FileController struct{}

// Upload maneja la solicitud de carga de archivos
func (f *FileController) Upload(c *gin.Context) {
	// Lógica de carga de archivos
	c.String(200, "Archivo cargado exitosamente")
}

// List maneja la solicitud de lista de archivos
func (f *FileController) List(c *gin.Context) {
	// Lógica para listar archivos
	c.String(200, "Lista de archivos")
}
