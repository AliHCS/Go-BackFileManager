package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// FileService maneja la lógica relacionada con archivos.
type FileService struct {
	client *mongo.Client // Cliente de MongoDB
}

// NewAuthService crea una nueva instancia de AuthService
func NewFilesService(client *mongo.Client) *FileService {
	return &FileService{client: client}
}

// UploadFile maneja la carga de un archivo y lo guarda localmente.
func (fs *FileService) UploadFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	// Crear un nombre único para el archivo
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), header.Filename)

	// Definir la ruta donde se almacenará el archivo
	filePath := filepath.Join("uploads", fileName)

	// Crear el directorio si no existe
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return "", fmt.Errorf("no se pudo crear el directorio: %v", err)
	}

	// Crear el archivo en la ruta definida
	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("no se pudo crear el archivo: %v", err)
	}
	defer out.Close()

	// Copiar el contenido del archivo subido al archivo local
	if _, err = io.Copy(out, file); err != nil {
		return "", fmt.Errorf("error al guardar el archivo: %v", err)
	}

	return filePath, nil
}
