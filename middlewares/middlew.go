package middlewares

import (
	"net/http"

	"webservice/bd"

	"github.com/gin-gonic/gin"
)

/*ChequeoBD es el middlew que me permite conocer el estado de la BD */
func ChequeoBD() gin.HandlerFunc {
	return func(c *gin.Context) {
		if bd.CheckConnect() == 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Conexión perdida con la Base de Datos",
			})
			return
		}
		c.Next()
	}
}
