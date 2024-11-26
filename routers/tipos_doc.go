package routers

import (
	"net/http"
	"webservice/bd"

	"github.com/gin-gonic/gin"
)

func Get_Tipos(c *gin.Context) {
	tipos, err := bd.GetTipos()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No hay docentes"})
		return
	}

	c.JSON(http.StatusOK, tipos)
}
