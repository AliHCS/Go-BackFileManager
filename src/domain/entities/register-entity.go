package entities

import "errors"

// RegisterUserEntity representa la entidad del usuario en Go
type RegisterUserEntity struct {
	ID       string
	Name     string
	Email    string
	Password string
}

// NewRegisterUserEntity crea una nueva instancia de RegisterUserEntity a partir de un mapa
func NewRegisterUserEntity(data map[string]string) (*RegisterUserEntity, error) {
	// Extraer los valores del mapa
	id, okID := data["id"]
	name, okName := data["name"]
	email, okEmail := data["email"]
	password, okPassword := data["password"]

	// Validaciones similares a las del código en Node.js
	if !okID || id == "" {
		return nil, errors.New("missing id")
	}
	if !okName || name == "" {
		return nil, errors.New("missing name")
	}
	if !okEmail || email == "" {
		return nil, errors.New("missing email")
	}
	if !okPassword || password == "" {
		return nil, errors.New("missing password")
	}

	// Si todo es válido, se retorna la nueva entidad de usuario
	return &RegisterUserEntity{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
