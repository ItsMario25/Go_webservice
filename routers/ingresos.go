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
	valid, nombre, rol, err := bd.ValidarUsuario(idUser, password)

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
		tks := models.Tk{
			Token: tokenString,
			Rols:  rol,
		}
		c.JSON(200, tks)

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

	user, client, rol, err := jwt.DecodeJWT(tokenString)

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
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Permisos invalidos", "rol_user": rol})
		}
	}
}

func Validar_Token(c *gin.Context) {
	var input struct {
		Token string `json:"token"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	log.Println(input.Token)

	err := bd.Verificar_token_correo(input.Token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Token validado exitosamente"})
}
