package services

import (
	"FileManager/src/domain/dtos/auth"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// AuthService maneja la lógica de autenticación
type AuthService struct {
	// Puedes añadir más campos si es necesario, como servicios externos
}

// NewAuthService crea una nueva instancia de AuthService
func NewAuthService() *AuthService {
	return &AuthService{}
}

// RegisterUser registra un nuevo usuario
func (a *AuthService) RegisterUser(registerDto *auth.RegisterDTO) (string, error) {
	// Lógica simulada: Aquí puedes manejar la lógica de verificación de usuarios en la memoria
	// Simulación de que ya existe un usuario
	existingEmail := "example@example.com"
	if registerDto.Email == existingEmail {
		return "", errors.New("email already exists")
	}

	// Hashea la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing password")
	}
	registerDto.Password = string(hashedPassword)

	// Simulación de la creación de un nuevo usuario

	return registerDto.Password, nil
}
