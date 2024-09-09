package routers

import (
	"net/http"
	"webservice/bd"

	"github.com/gin-gonic/gin"
)

func Curso_facultad(c *gin.Context) {
	cursos, err := bd.GetCursos_facultad()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No hay periodo activo"})
		return
	}

	c.JSON(http.StatusOK, cursos)
}

func Get_Curso(c *gin.Context) {
	id := c.Param("id")

	cur, err := bd.GetCurso(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No hay curso"})
		return
	}

	c.JSON(http.StatusOK, cur)
}
