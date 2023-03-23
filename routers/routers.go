package routers

import (
	"net/http"

	"webservice/models"

	"github.com/gin-gonic/gin"
)

/*Registro es la funcion para crear en la BD el registro de usuario */
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

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
	})
}
