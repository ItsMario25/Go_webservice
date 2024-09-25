package routers

import (
	"net/http"
	"webservice/bd"

	"github.com/gin-gonic/gin"
)

func Curso_facultad(c *gin.Context) {
	cursos, err := bd.GetCursos()

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

func Get_Cursos_asignados(c *gin.Context) {
	materias_asignadas, err := bd.GetCursosAsignados()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los cursos asignados"})
	}
	c.JSON(http.StatusOK, gin.H{"cursos_asignados": materias_asignadas})
}
