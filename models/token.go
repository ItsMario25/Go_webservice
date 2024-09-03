package models

import "github.com/dgrijalva/jwt-go"

// Claims define la estructura para almacenar los datos en el JWT
type Claims struct {
	Username string `json:"username"`
	ClientID string `json:"client_id"`
	RolUser  string `json:"rol_user"`
	jwt.StandardClaims
}

type TokenRequest struct {
	ClientID string `json:"client_id"`
	Rol_us   string `json:"rol"`
}

type Credentials struct {
	Usuario    string `json:"usuario"`
	Contrasena string `json:"contrasena"`
	ClientID   string `json:"client_id"`
}
