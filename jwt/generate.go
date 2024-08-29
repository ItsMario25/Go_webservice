package jwt

import (
	"time"
	"webservice/models"

	"github.com/dgrijalva/jwt-go"
)

func GenerarToken(nombre string, userID string, jwtKey string, rol string) (string, error) {

	var jwtk = []byte(jwtKey)
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &models.Claims{
		Username: nombre,
		ClientID: userID,
		RolUser:  rol,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtk)

	return tokenString, err
}
