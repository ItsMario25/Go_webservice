package routers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webservice/bd"
	"webservice/models"

	"github.com/gin-gonic/gin"
)

func GetPeriodos(c *gin.Context) {
	consulta, err := bd.GetPeriod()

	if err != nil {
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, consulta)
	}
}

func CargarPeriodo(c *gin.Context) {
	var periodo models.PeriodoAcademico

	if err := c.BindJSON(&periodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar si los campos no están vacíos
	if periodo.Periodo == "" || periodo.Inicio == "" || periodo.Fin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos incompletos"})
		return
	}

	num, err := strconv.Atoi(periodo.Periodo)

	if err != nil {
		fmt.Println("Error:", err)
	}

	err = bd.SetPeriod(num, periodo.Inicio, periodo.Fin)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fallo de insercion"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Periodo académico cargado con éxito"})

	}
}
