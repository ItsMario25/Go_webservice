package utilidades

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncriptarPassword(pass string) (string, error) {
	costo := 9
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}

func ValidarPassword(pass string, hashpass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(pass))
	if err != nil {
		log.Println("La contraseña no coincide")
		return false
	} else {
		log.Println("La contraseña coincide")
		return true
	}
}
