package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ValidarExpiracion(tokenString string, secretKey []byte) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return fmt.Errorf("error parsing token: %v", err)
	}

	// Verificamos si el token es válido
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Verificamos la expiración
		expiration, ok := claims["exp"].(float64)
		if !ok {
			return fmt.Errorf("token does not contain an exp claim")
		}

		// Convertimos el tiempo de expiración en un objeto Time
		expirationTime := time.Unix(int64(expiration), 0)
		if time.Now().After(expirationTime) {
			return fmt.Errorf("token has expired")
		}

		// Token es válido y no ha expirado
		fmt.Println("Token is valid and has not expired")
		return nil
	}

	return fmt.Errorf("invalid token")
}
