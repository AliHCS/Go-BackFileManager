package files

import (
	"errors"
	"time"
)

// UploadFileDto representa el DTO para la carga de archivos
type UploadFileDto struct {
	UserID     string     // ID del usuario al que pertenece el archivo
	File       FileInfo   // Información del archivo
	Filename   string     // Nombre del archivo
	Path       string     // Ruta donde se almacena el archivo
	MimeType   string     // Tipo de archivo (ej. image/png, application/pdf)
	Size       int64      // Tamaño del archivo en bytes
	UploadedAt time.Time  // Fecha y hora de subida
	UpdatedAt  *time.Time // Fecha y hora de la última actualización (opcional)
}

// FileInfo contiene información sobre el archivo
type FileInfo struct {
	OriginalName string // Nombre original del archivo
	Filename     string // Nombre del archivo en el servidor
	Path         string // Ruta donde se almacena el archivo
	MimeType     string // Tipo de archivo
	Size         int64  // Tamaño del archivo en bytes
}

// NewUploadFileDto crea una nueva instancia de UploadFileDto
func NewUploadFileDto(userId string, file FileInfo) (*UploadFileDto, error) {
	/* if userId <= 0 {
		return nil, errors.New("el userId debe ser un número válido")
	} */
	if file.OriginalName == "" || file.Filename == "" || file.Path == "" || file.MimeType == "" || file.Size <= 0 {
		return nil, errors.New("el archivo es requerido y debe contener información válida")
	}

	uploadedAt := time.Now() // Fecha actual como `uploadedAt`
	return &UploadFileDto{
		UserID:     userId,
		File:       file,
		Filename:   file.Filename,
		Path:       file.Path,
		MimeType:   file.MimeType,
		Size:       file.Size,
		UploadedAt: uploadedAt,
		UpdatedAt:  nil, // Puedes asignar un valor aquí si lo necesitas
	}, nil
}
