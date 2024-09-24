package auth

import (
	"errors"
	"regexp"
)

// RegisterDTO es el equivalente a la clase RegisterDto en TypeScript
type RegisterDTO struct {
	Email    string
	Password string
	Name     string
}

// ValidateEmail valida si el email tiene un formato correcto
func ValidateEmail(email string) bool {
	// Regex básica para validar un correo electrónico
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

// NewRegisterDTO
func NewRegisterDTO(props map[string]string) (*RegisterDTO, error) {
	email, emailOk := props["email"]
	password, passwordOk := props["password"]
	name, nameOk := props["name"]

	if !emailOk || email == "" {
		return nil, errors.New("el Email debe ser necesario")
	}
	if !nameOk || name == "" {
		return nil, errors.New("el name debe ser necesario")
	}
	if !passwordOk || password == "" {
		return nil, errors.New("el password debe ser necesario")
	}
	if len(password) < 6 {
		return nil, errors.New("password too short")
	}
	if !ValidateEmail(email) {
		return nil, errors.New("el formato del email es inválido")
	}

	return &RegisterDTO{
		Email:    email,
		Password: password,
		Name:     name,
	}, nil
}

// Values retorna un mapa con los valores del DTO
func (dto *RegisterDTO) Values() map[string]interface{} {
	returnObj := make(map[string]interface{})

	if dto.Email != "" {
		returnObj["email"] = dto.Email
	}
	if dto.Password != "" {
		returnObj["password"] = dto.Password
	}
	if dto.Name != "" {
		returnObj["name"] = dto.Name
	}

	return returnObj
}
