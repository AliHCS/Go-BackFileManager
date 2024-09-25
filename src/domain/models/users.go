package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`     // ID en BSON y JSON
	Name      string             `bson:"name" json:"name"`            // Nombre del usuario
	Email     string             `bson:"email" json:"email_address"`  // Correo en JSON como email_address
	Password  string             `bson:"password" json:"password"`    // Contraseña
	CreatedAt time.Time          `bson:"createdAt" json:"created_at"` // Fecha de creación
	UpdatedAt primitive.DateTime `bson:"updatedAt" json:"updated_at"` // Fecha de actualización
}
