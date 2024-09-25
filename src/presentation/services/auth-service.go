package services

import (
	"FileManager/src/config"
	"FileManager/src/domain/dtos/auth"
	"FileManager/src/domain/models"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AuthService maneja la lógica de autenticación
type AuthService struct {
	client *mongo.Client // Cliente de MongoDB
	// Puedes añadir más campos si es necesario, como servicios externos
}

// NewAuthService crea una nueva instancia de AuthService
func NewAuthService(client *mongo.Client) *AuthService {
	return &AuthService{client: client}
}

// RegisterUser registra un nuevo usuario
func (a *AuthService) RegisterUser(registerDto *auth.RegisterDTO) (string, error) {
	// Verificar si el email ya existe en la base de datos
	collection := a.client.Database(config.LoadEnv().MONGO_BD_NAME).Collection("users") // Reemplaza con tu base de datos y colección
	log.Printf("Intentando registrar usuario: %v", registerDto)
	// Hacer la consulta para verificar si el correo ya existe
	var existingUser models.User
	err := collection.FindOne(context.Background(), bson.M{"email": registerDto.Email}).Decode(&existingUser)

	// Si hay un error que no es mongo.ErrNoDocuments, significa que ocurrió un error
	if err != nil && err != mongo.ErrNoDocuments {
		return "", err
	}

	// Si existingUser tiene valores, significa que el usuario ya existe
	if existingUser.Email != "" {
		return "", errors.New("el correo electrónico ya está registrado")
	}

	// Hashea la contraseña usando la función extraída
	hashedPassword, err := config.HashPassword(registerDto.Password)
	if err != nil {
		return "", err
	}
	registerDto.Password = string(hashedPassword)

	// Crear un nuevo usuario en la base de datos
	newUser := models.User{
		Email:    registerDto.Email,
		Name:     registerDto.Name,
		Password: registerDto.Password, // O almacenar el hashedPassword
	}

	// Insertar el nuevo usuario en la colección
	_, err = collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return "", err
	}

	return "Usuario registrado con éxito", nil
}
