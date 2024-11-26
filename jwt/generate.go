package jwt

import (
	"log"
	"time"
	"webservice/models/core"
	"webservice/utilities"

	"github.com/dgrijalva/jwt-go"
)

func GenerarToken(nombre string, userID string, jwtKey string, rol string) (string, error) {

	pass, err := utilities.EncriptarPassword(jwtKey)

	if err != nil {
		log.Println("Error al encriptar en la generacion del token")
	}
	var jwtk = []byte(pass)
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &core.Claims{
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
