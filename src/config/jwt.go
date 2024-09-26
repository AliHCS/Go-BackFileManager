package config

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(LoadEnv().SECRET_KEY) // Cambia esto por una clave más segura

// Claims es la estructura para los claims del token
type Claims struct {
	Email                string `json:"email"`
	ID                   string `json:"id"`
	jwt.RegisteredClaims        // Usa la nueva estructura para los claims registrados
}

// GenerateToken genera un nuevo token JWT
func GenerateToken(email, userID string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour) // Establece el tiempo de expiración

	claims := &Claims{
		Email: email,
		ID:    userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Establece el tiempo de expiración
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// VerifyToken verifica si el token JWT es válido
func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err // Retorna el error si no se pudo analizar el token
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid // Retorna un error si el token no es válido
	}

	return claims, nil // Retorna los claims si el token es válido
}
