package routers

import (
	"log"
	"net/http"
	"webservice/bd"
	"webservice/jwt"

	"github.com/gin-gonic/gin"
)

func Get_programa(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	user, _, _, err := jwt.DecodeJWT(tokenString)

	if err != nil {
		log.Println("token no descifrado")
	}

	us, err := bd.GetUsuarioid(user)

	if err != nil {
		log.Println("Usuario no encontrado")
	}

	miembro, err := bd.Get_Secretario(us)

	if err != nil {
		log.Println("Miembro del consejo no encontrado")
	}

	log.Println(miembro.IDFacultad)

	pr, err := bd.GetProgramasPorFacultad(miembro.IDFacultad)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error al traer programas"})
	} else {
		c.JSON(http.StatusOK, pr)
	}
}
