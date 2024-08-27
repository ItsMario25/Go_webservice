package routers

import (
	"log"
	"net/http"
	"time"

	"webservice/bd"
	"webservice/models"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your_secret_key")

func ValidarLogin(c *gin.Context) {

	var credentials struct {
		Usuario    string `json:"usuario"`
		Contrasena string `json:"contrasena"`
		ClientID   string `json:"client_id"`
	}

	var estudiante models.Estudiante

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	log.Printf("Received data: Usuario=%s, Contrasena=%s, ClientID=%s", credentials.Usuario, credentials.Contrasena, credentials.ClientID)

	idUser := credentials.Usuario
	password := credentials.Contrasena

	// Validación de credenciales
	// Buscar el usuario y su contraseña
	if err := bd.DB.Table("estudiante").Select("codigo_estudiante, clave_estudiante, id_user").
		Where("clave_estudiante = ? AND codigo_estudiante = ?", password, idUser).
		First(&estudiante).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &models.Claims{
		Username: credentials.Usuario,
		ClientID: credentials.ClientID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})

}
