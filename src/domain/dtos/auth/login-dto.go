package auth

import (
	"FileManager/src/domain/validations"
	"errors"
)

// LoginDTO es el Data Transfer Object para el login
type LoginDTO struct {
	Email    string
	Password string
}

// NewLoginDTO valida los datos y devuelve un nuevo LoginDTO si todo es correcto
func NewLoginDTO(props map[string]string) (*LoginDTO, error) {
	email, emailOk := props["email"]
	password, passwordOk := props["password"]
	if !emailOk || email == "" {
		return nil, errors.New("el Email debe ser necesario")
	}

	if !passwordOk || password == "" {
		return nil, errors.New("el password debe ser necesario")
	}

	if !validations.ValidateEmail(email) {
		return nil, errors.New("el formato del email es inválido")
	}
	// Si todas las validaciones pasan, se crea y devuelve el DTO
	return &LoginDTO{
		Email:    email,
		Password: password,
	}, nil
}

// Values retorna un mapa con los valores del DTO
func (dto *LoginDTO) Values() map[string]interface{} {
	returnObj := make(map[string]interface{})

	if dto.Email != "" {
		returnObj["email"] = dto.Email
	}
	if dto.Password != "" {
		returnObj["password"] = dto.Password
	}

	return returnObj
}
