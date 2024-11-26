package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

func DecodeJWT(tokenString string) (string, string, string, error) {
	// Parsear el token sin validarlo
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", "", "", err
	}

	// Acceder a las claims (afirmaciones) del token
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		us := claims["username"].(string)
		cl := claims["client_id"].(string)
		rol := claims["rol_user"].(string)

		return us, cl, rol, nil
	} else {
		return "", "", "", err
	}
}
