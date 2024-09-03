package routers

import (
	"log"
	"net/http"

	"webservice/bd"
	"webservice/jwt"
	"webservice/models"

	"github.com/gin-gonic/gin"
)

func ValidarLogin(c *gin.Context) {

	var credenciales models.Credentials

	if err := c.BindJSON(&credenciales); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	log.Printf("Received data: Usuario=%s, Contrasena=%s, ClientID=%s", credenciales.Usuario, credenciales.Contrasena, credenciales.ClientID)

	idUser := credenciales.Usuario
	password := credenciales.Contrasena

	// Validación de credenciales
	valid, nombre, rol, err := bd.ValidarUsuario(bd.DB, idUser, password)

	if err != nil {
		c.JSON(500, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if valid {

		tokenString, err := jwt.GenerarToken(nombre, credenciales.ClientID, password, rol)
		if err != nil {
			c.JSON(500, gin.H{"error": "No se pudo generar Token"})
			return
		}
		c.JSON(200, gin.H{"token": tokenString})

	} else {
		c.JSON(401, gin.H{"error": "Credenciales incorrectas"})
	}

}

func ValidarToken(c *gin.Context) {
	var req models.TokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	log.Println(tokenString)
	user, client, rol, err := jwt.DecodeJWT(tokenString)
	log.Println(req.ClientID)
	log.Println(req.Rol_us)
	if err != nil {
		log.Println("Error al decodificar token : ", err)
	} else {
		if req.ClientID == client && req.Rol_us == rol {
			log.Println("Id de sesion y rol validos")
			clv, err := bd.Get_clave(user, rol)
			if err == nil {
				err = jwt.ValidarExpiracion(tokenString, []byte(clv))
				if err != nil {
					c.JSON(http.StatusOK, gin.H{"message": "Token válido", "rol_user": rol})
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
				}
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Permisos invalidos"})
		}
	}
}
