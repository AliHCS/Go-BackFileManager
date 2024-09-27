package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// File representa un archivo subido por un usuario
type File struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`       // ID del archivo
	UserID     primitive.ObjectID `bson:"userId" json:"user_id"`         // ID del usuario al que pertenece el archivo
	FileName   string             `bson:"fileName" json:"file_name"`     // Nombre del archivo
	Path       string             `bson:"path" json:"path"`              // Ruta del archivo en el servidor
	MimeType   string             `bson:"mimeType" json:"mime_type"`     // Tipo MIME del archivo
	Size       int64              `bson:"size" json:"size"`              // Tamaño del archivo en bytes
	UploadedAt time.Time          `bson:"uploadedAt" json:"uploaded_at"` // Fecha y hora de subida
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updated_at"`   // Fecha y hora de la última actualización
}
