package routers

import (
	"fmt"
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

	// Aquí podrías guardar las respuestas en la base de datos, por ejemplo.
	fmt.Printf("Nombre del curso: %s\n", evaluacion.NombreCurso)
	fmt.Printf("Nombre del docente: %s\n", evaluacion.NombreDocente)
	fmt.Printf("Respuestas: %+v\n", evaluacion.Respuestas)

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Evaluación recibida correctamente"})
}
