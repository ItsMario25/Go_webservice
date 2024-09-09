package routers

import (
	"net/http"
	"webservice/bd"

	"github.com/gin-gonic/gin"
)

func Get_Docentes(c *gin.Context) {
	docentes, err := bd.GetDocentes()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No hay docentes"})
		return
	}

	c.JSON(http.StatusOK, docentes)
}
