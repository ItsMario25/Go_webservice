package routers

import (
	"fmt"
	"net/http"
	"strconv"
	"webservice/bd"
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
