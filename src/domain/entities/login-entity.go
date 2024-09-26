package entities

import "errors"

// LoginUserEntity representa la entidad de usuario al iniciar sesión
type LoginUserEntity struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"` // Si deseas almacenar el token en la entidad
}

// NewLoginUserEntity crea una nueva entidad de usuario al iniciar sesión
func NewLoginUserEntity(data map[string]string) (*LoginUserEntity, error) {
	id, ok := data["id"]
	if !ok {
		return nil, errors.New("el ID es necesario")
	}
	name, ok := data["name"]
	if !ok {
		return nil, errors.New("el nombre es necesario")
	}
	email, ok := data["email"]
	if !ok {
		return nil, errors.New("el email es necesario")
	}
	token, ok := data["token"] // Si decides incluir el token
	if !ok {
		return nil, errors.New("el token es necesario")
	}

	return &LoginUserEntity{
		ID:    id,
		Name:  name,
		Email: email,
		Token: token,
	}, nil
}
