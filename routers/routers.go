package routers

import (
	"log"

	"webservice/bd"
	"webservice/jwt"

	"github.com/gin-gonic/gin"
)

func ValidarLogin(c *gin.Context) {

	var credentials struct {
		Usuario    string `json:"usuario"`
		Contrasena string `json:"contrasena"`
		ClientID   string `json:"client_id"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	log.Printf("Received data: Usuario=%s, Contrasena=%s, ClientID=%s", credentials.Usuario, credentials.Contrasena, credentials.ClientID)

	idUser := credentials.Usuario
	password := credentials.Contrasena

	// Validación de credenciales
	valid, nombre, rol, err := bd.ValidarUsuario(bd.DB, idUser, password)

	if err != nil {
		c.JSON(500, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if valid {

		tokenString, err := jwt.GenerarToken(nombre, credentials.ClientID, password, rol)
		if err != nil {
			c.JSON(500, gin.H{"error": "No se pudo generar Token"})
			return
		}
		c.JSON(200, gin.H{"token": tokenString})

	} else {
		c.JSON(401, gin.H{"error": "Credenciales incorrectas"})
	}

}
