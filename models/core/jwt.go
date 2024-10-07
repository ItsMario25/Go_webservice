package core

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	ClientID string `json:"client_id"`
	RolUser  string `json:"rol_user"`
	jwt.StandardClaims
}

type Tk struct {
	Token string `json:"token"`
	Rols  string `json:"rol"`
}
