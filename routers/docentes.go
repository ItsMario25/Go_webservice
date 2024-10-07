package routers

import (
	"log"
	"net/http"
	"strconv"
	"webservice/bd"
	"webservice/jwt"
	"webservice/models/request"

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
	var asignar request.GetAsignacion

	if err := c.BindJSON(&asignar); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	docen, err1 := strconv.Atoi(asignar.IDdocente)
	curse, err2 := strconv.Atoi(asignar.Idcurso)
	tipo := asignar.IDtipo

	if err1 != nil {
		log.Println("Error al convertir str1:", err1)
	}
	if err2 != nil {
		log.Println("Error al convertir str2:", err2)
	}

	log.Println(docen)
	log.Println(curse)
	log.Println(tipo)
	err := bd.SetEjerciendo(docen, curse, tipo)

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

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Estudiante no encontrado"})
	}
	resultados, err := bd.GetDocentesActuales(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en la consulta"})
	}

	c.JSON(http.StatusOK, resultados)
}

func Get_Docentes_curso(c *gin.Context) {
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

	id, err := bd.GetDocente(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Estudiante no encontrado"})
	}
	resultados, err := bd.GetEjerciendoActual(id.IDDocente)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en la consulta"})
	}

	c.JSON(http.StatusOK, resultados)
}

func Get_Docentes_facultad(c *gin.Context) {
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

	consejo, err := bd.Get_Consejo(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Miembro del consejo no encontrado"})
	}

	resultados, err := bd.Get_Docentes_facultad(consejo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Docentes no encontrados"})
	}

	c.JSON(http.StatusOK, resultados)

}

func ValidarEjerciendo(c *gin.Context) {
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

	periodo, err := bd.GetPeriodoActivo()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Periodo de evaluacion no activo"})
	}

	cursos, err := bd.Get_materias_evaluadas(user, periodo.IDPeriodoEvl)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No hay evaluacion realizada"})
	}

	if len(cursos) > 0 {
		ejer := false
		c.JSON(http.StatusBadRequest, ejer)
	} else {
		ejer := true
		c.JSON(http.StatusOK, ejer)
	}
}

func Get_docentes_evaluados(c *gin.Context) {
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

	periodo, err := bd.GetPeriodoActivo()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Periodo de evaluacion no activo"})
	}

	cursos, err := bd.Get_docentes_evl(user, periodo.IDPeriodoEvl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener informacion de materias evaluadas"})
	} else {
		c.JSON(http.StatusOK, gin.H{"docentes_evaluados": cursos})
	}

}
