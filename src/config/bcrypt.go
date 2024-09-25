package config

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashea una contraseña usando bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing password")
	}
	return string(hashedPassword), nil
}

// VerifyPassword verifica si la contraseña ingresada coincide con la contraseña hasheada
func VerifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("contraseña incorrecta")
	}
	return nil
}
