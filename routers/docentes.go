package routers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webservice/bd"
	"webservice/jwt"
	"webservice/models"

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

func AsignarEjerciendo(c *gin.Context) {
	var asignar models.GetAsignacion

	if err := c.BindJSON(&asignar); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	docen, err1 := strconv.Atoi(asignar.IDdocente)
	curse, err2 := strconv.Atoi(asignar.Idcurso)

	if err1 != nil {
		fmt.Println("Error al convertir str1:", err1)
	}
	if err2 != nil {
		fmt.Println("Error al convertir str2:", err2)
	}

	err := bd.SetEjerciendo(docen, curse)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingreso Fallido"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Ejerciendo asignado"})
	}
}

func Get_Docentes_materia(c *gin.Context) {
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

	id, err := bd.GetEstudianteid(user)

	log.Println(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Estudiante no encontrado"})
	}
	resultados, err := bd.GetDocentesActuales(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en la consulta"})
	}

	c.JSON(http.StatusOK, resultados)
}
