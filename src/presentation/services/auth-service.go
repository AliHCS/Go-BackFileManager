package services

import (
	"FileManager/src/config"
	"FileManager/src/domain/dtos/auth"
	"FileManager/src/domain/entities"
	"FileManager/src/domain/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterUserResponse es la respuesta sin la contraseña
type RegisterUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// LoginResponse es la respuesta para el login exitoso
type LoginResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

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
func (a *AuthService) RegisterUser(registerDto *auth.RegisterDTO) (*RegisterUserResponse, error) {
	// Verificar si el email ya existe en la base de datos
	collection := a.client.Database(config.LoadEnv().MONGO_BD_NAME).Collection("users") // Reemplaza con tu base de datos y colección
	// Hacer la consulta para verificar si el correo ya existe
	var existingUser models.User
	err := collection.FindOne(context.Background(), bson.M{"email": registerDto.Email}).Decode(&existingUser)

	// Si hay un error que no es mongo.ErrNoDocuments, significa que ocurrió un error
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	// Si existingUser tiene valores, significa que el usuario ya existe
	if existingUser.Email != "" {
		return nil, errors.New("el correo electrónico ya está registrado")
	}

	// Hashea la contraseña usando la función extraída
	hashedPassword, err := config.HashPassword(registerDto.Password)
	if err != nil {
		return nil, err
	}
	registerDto.Password = string(hashedPassword)

	// Crear un nuevo usuario en la base de datos
	// Insertar el nuevo usuario en la colección
	newUser := models.User{
		Email:    registerDto.Email,
		Name:     registerDto.Name,
		Password: registerDto.Password, // O almacenar el hashedPassword
	}
	result, err := collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return nil, err
	}
	// Obtener el ID del nuevo usuario insertado
	newUserID := result.InsertedID.(primitive.ObjectID).Hex() // Asegúrate de hacer la conversión adecuada

	// Crear un mapa con los datos del nuevo usuario
	userData := map[string]string{
		"id":       newUserID,
		"name":     registerDto.Name,
		"email":    registerDto.Email,
		"password": registerDto.Password,
	}
	// Crear la entidad del nuevo usuario
	newUserEntity, err := entities.NewRegisterUserEntity(userData)
	if err != nil {
		return nil, err
	}
	// Crear la respuesta sin la contraseña
	response := &RegisterUserResponse{
		ID:    newUserEntity.ID,
		Name:  newUserEntity.Name,
		Email: newUserEntity.Email,
	}
	return response, nil
}

// LoginUser realiza la autenticación de un usuario
func (a *AuthService) LoginUser(loginDto *auth.LoginDTO) (*LoginResponse, error) {
	// Verificar si el usuario existe
	collection := a.client.Database(config.LoadEnv().MONGO_BD_NAME).Collection("users")

	var existingUser models.User

	err := collection.FindOne(context.Background(), bson.M{"email": loginDto.Email}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}

	// Verificar si la contraseña es correcta
	err = config.VerifyPassword(existingUser.Password, loginDto.Password)
	if err != nil {
		return nil, err // Retorna el error de contraseña incorrecta
	}

	// Generar el token JWT
	token, err := config.GenerateToken(existingUser.Email, existingUser.ID.Hex())
	if err != nil {
		return nil, err // Retorna el error si no se pudo generar el token
	}

	// Retornar respuesta de éxito en el login
	return &LoginResponse{
		ID:    existingUser.ID.Hex(), // Convertir el ObjectID a string
		Name:  existingUser.Name,
		Email: existingUser.Email,
		Token: token, // Incluir el token en la respuesta
	}, nil
}
