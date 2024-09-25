package routers

import (
	"log"
	"net/http"
	"webservice/bd"
	"webservice/jwt"

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

func Get_cursos_evaluados(c *gin.Context) {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "token no descifrado"})
	}

	log.Println(user)

	periodo, err := bd.GetPeriodoActivo()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Periodo de evaluacion no activo"})
	}

	log.Println(periodo)

	bd.Get_materias_evaluadas(user, periodo.IDPeriodoEvl)

}
