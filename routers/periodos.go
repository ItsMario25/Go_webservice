package routers

import (
	"log"
	"net/http"
	"webservice/bd"
	"webservice/jwt"
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
	var periodo models.PeriodoAc

	if err := c.BindJSON(&periodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar si los campos no están vacíos
	if periodo.Periodo == "" || periodo.Inicio == "" || periodo.Fin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos incompletos"})
		return
	}

	err := bd.SetPeriod(periodo.Periodo, periodo.Inicio, periodo.Fin)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fallo de insercion"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Periodo académico cargado con éxito"})
	}
}

func PeriodoActivo(c *gin.Context) {

	periodoact, err := bd.GetPeriodoActivo()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No hay periodo activo"})
		return
	}

	c.JSON(http.StatusOK, periodoact)

}

func PeriodoACactivo(c *gin.Context) {

	peridoacc, err := bd.GetPeriodoAcActivo()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No hay periodo activo"})
		return
	}

	c.JSON(http.StatusOK, peridoacc)
}

func Editarperiodo(c *gin.Context) {
	var periodo models.PeriodoAc

	// Obtener el ID del periodo desde la URL
	id := c.Param("id")

	// Obtener datos de la solicitud JSON
	if err := c.BindJSON(&periodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Validar que el ID no esté vacío
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de periodo no proporcionado"})
		return
	}

	err := bd.PutPeriodo(id, periodo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al actualizar el periodo"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Periodo actualizado con éxito"})
	}

}

func Get_Periodos_facultad(c *gin.Context) {
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
		log.Println("token no descifrado")
	}

	us, err := bd.GetUsuarioid(user)

	if err != nil {
		log.Println("Usuario no encontrado")
	}
	log.Println(us)

	sec, err := bd.Get_Secretario(us)

	if err != nil {
		log.Println("Secretario no encontrado")
	}

	periodos, err := bd.GetPeriodo_facultad(sec.IDFacultad)

	if err != nil {
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, periodos)
	}
}
