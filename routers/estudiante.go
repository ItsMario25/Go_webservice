package routers

import (
	"net/http"
	"webservice/bd"

	"github.com/gin-gonic/gin"
)

func Get_Criterios_estudiante(c *gin.Context) {
	format := "estudiante"

	criterios, err := bd.Get_criterios_bd(format)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Criterios no encontrados"})
	} else {
		c.JSON(http.StatusOK, criterios)
	}
}
