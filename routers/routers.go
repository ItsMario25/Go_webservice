package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Usuario    string `json:"usuario"`
	Contrasena string `json:"contrasena"`
}

/*Registro es la funcion para crear en la BD el registro de usuario
func Registro(c *gin.Context) {
	var t models.Usuario

	if err := c.Bind(&t); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error en los datos recibidos " + err.Error(),
		})
		return
	}

	if len(t.Email) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "El email de usuario es requerido",
		})
		return
	}

	if len(t.Password) < 6 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Debe especificar una contraseña de al menos 6 caracteres",
		})
		return
	}

	GetIndex(c)
}*/

func ValidarLogin(c *gin.Context) {
	var data LoginData

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Aquí puedes manejar la verificación de usuario y contraseña.
	// Por ahora, simplemente retornamos los datos que recibimos.
	c.JSON(http.StatusOK, gin.H{
		"message": "Inicio de sesión exitoso",
		"usuario": data.Usuario,
	})

}
