package routers

import (
	"net/http"
	"webservice/bd"
	"webservice/models"

	"github.com/gin-gonic/gin"
)

func Guardar_autoevaluacion(c *gin.Context) {
	var autoevaluacion models.EvaluacionDocente

	if err := c.ShouldBindJSON(&autoevaluacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bd.Guardar_evl_docente(autoevaluacion)

	c.JSON(http.StatusOK, gin.H{"message": "Evaluación recibida correctamente"})
}

func Guardar_evaluacion(c *gin.Context) {
	var evaluacion models.FormatoEvaluacion

	if err := c.ShouldBindJSON(&evaluacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bd.Guardar_evl(evaluacion)

	c.JSON(http.StatusOK, gin.H{"message": "Evaluación recibida correctamente"})
}

func Guardar_evaluacion_facultad(c *gin.Context) {
	var evaluacion models.FormatoEvaluacion

	if err := c.ShouldBindJSON(&evaluacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bd.Guardar_evl_facultad(evaluacion)

	c.JSON(http.StatusOK, gin.H{"message": "Evaluación recibida correctamente"})
}
