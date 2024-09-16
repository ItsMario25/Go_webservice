package routers

import (
	"net/http"
	"webservice/bd"
	"webservice/models"

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

func Guardar_evaluacion(c *gin.Context) {
	var evaluacion models.EvaluacionEstudiante

	// Bind JSON a la estructura EvaluacionEstudiante
	if err := c.ShouldBindJSON(&evaluacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bd.Guardar_evl(evaluacion)

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Evaluación recibida correctamente"})
}
