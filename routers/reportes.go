package routers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"webservice/bd"
	"webservice/jwt"
	"webservice/models/request"
	"webservice/utilities"

	"github.com/gin-gonic/gin"
)

func Historial_individual(c *gin.Context) {
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
		log.Println("Error en la desencriptacion")
	}

	doc, err := bd.GetDocente(user)

	if err != nil {
		log.Println("Error en la desencriptacion")
	}

	ejer, err := bd.Get_reportes_semestre(doc.IDDocente)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener el reporte de evaluaciones"})
	}
	for i, reporte := range ejer {
		ejer[i].PeriodoAcad = strings.TrimPrefix(reporte.PeriodoAcad, "PA")
	}

	c.JSON(200, ejer)
}

func Reporte_individual(c *gin.Context) {
	var request request.ReporteRequest

	// Vincular el JSON recibido a la estructura `ReporteRequest`
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

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

	doc, err := bd.GetDocente(user)

	if err != nil {
		log.Println("Docente no encontrado")
	}

	periodoAcademico := request.PeriodoAcademico
	vincula := request.Vinculacion
	idPeriodoEvl := fmt.Sprintf("PE%s", periodoAcademico)
	cons, err := bd.Get_resultados_formato_consejo(idPeriodoEvl, us)

	if err != nil {
		log.Println("error en encontrar resultados de consejo")
	}

	est, err := bd.Get_resultados_formato_estudiante(idPeriodoEvl, us)

	if err != nil {
		log.Println("error en encontrar resultados de estudiante")
	}

	docs, err := bd.Get_resultados_formato_docente(idPeriodoEvl, us)

	if err != nil {
		log.Println("error en encontrar resultados de docente")
	}

	promedioConsejo := bd.CalcularCalificacionTotal(cons)
	promedioDocente := bd.CalcularCalificacionTotal(docs)
	promediosEstudiantes, estudianteCurso := bd.CalcularCalificacionPorCurso(est)

	data := gin.H{
		"Periodo":     idPeriodoEvl,
		"Profesor":    user,
		"Ide":         doc.IDDocente,
		"Vinculacion": vincula,
		"Consejo":     fmt.Sprintf("%.2f", promedioConsejo),
		"Docente":     fmt.Sprintf("%.2f", promedioDocente),
		"Estudiantes": promediosEstudiantes,
		"Cursos":      estudianteCurso,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8081/datos", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("error en la petición: %v", resp.Status)
	}

	// Leer el contenido del PDF desde la respuesta
	pdfBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error al leer el cuerpo de la respuesta:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el PDF"})
		return
	}

	config, err := bd.GetConfigSeguridad("copia_controlada")
	if err != nil {
		log.Println("Error al obtener configuración de seguridad:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la configuración"})
		return
	}

	if config.Estado {
		pdfHash := utilities.GeneratePDFHash(pdfBytes)

		// Guardar en la base de datos
		err = bd.SavePDFHash(user, pdfHash)
		if err != nil {
			log.Println("Error al guardar el hash:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el hash"})
			return
		}
	}

	// Enviar el PDF al front-end
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}

func Reporte_general(c *gin.Context) {
	var request request.ReportegRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

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

	sec, err := bd.Get_Secretario(us)

	if err != nil {
		log.Println("Secretario no encontrado")
	}

	log.Println(sec)

	idPeriodoEvl := fmt.Sprintf("PE%s", request.PeriodoAcademico)
	log.Println(idPeriodoEvl)
	log.Println(request.Programanombre)

	idprograma, err := bd.GerProgramaPorNombre(request.Programanombre)

	if err != nil {
		log.Println("Programa no encontrado")
	}

	log.Println(idprograma.IdPrograma)

	doc, err := bd.GetDocentesPorProgramaYPeriodo(idprograma.IdPrograma, idPeriodoEvl)

	if err != nil {
		log.Println("Docentes no encontrados")
	}

	log.Println(doc)

	var notasConsejo []string
	var notasAutoevalua []string
	var nombresCursos []string
	var idecursos []string
	var promediosCursos []string
	var estudian []int
	for _, docente := range doc {
		uss, err := bd.GetUsuarioid(docente.Nombre)

		if err != nil {
			log.Println("Usuario no encontrado")
		}

		cons, err := bd.Get_resultados_formato_consejo(idPeriodoEvl, uss)

		if err != nil {
			log.Println("error en encontrar resultados de consejo")
		}

		est, err := bd.Get_resultados_formato_estudiante(idPeriodoEvl, uss)

		if err != nil {
			log.Println("error en encontrar resultados de estudiante")
		}

		docs, err := bd.Get_resultados_formato_docente(idPeriodoEvl, uss)

		if err != nil {
			log.Println("error en encontrar resultados de docente")
		}

		promedioConsejo := bd.CalcularCalificacionTotal(cons)
		promedioDocente := bd.CalcularCalificacionTotal(docs)
		promediosEstudiantes, estudianteCurso := bd.CalcularCalificacionPorCurso(est)

		notasConsejo = append(notasConsejo, fmt.Sprintf("%.2f", promedioConsejo))
		notasAutoevalua = append(notasAutoevalua, fmt.Sprintf("%.2f", promedioDocente))

		for nombre, est := range estudianteCurso {
			nombresCursos = append(nombresCursos, nombre)
			estudian = append(estudian, est)
		}

		for codigo, nota := range promediosEstudiantes {
			idecursos = append(idecursos, codigo)
			promediosCursos = append(promediosCursos, nota)
		}

	}

	log.Println(notasAutoevalua)
	log.Println(notasConsejo)
	log.Println(nombresCursos)
	log.Println(idecursos)
	log.Println(promediosCursos)
	log.Println(estudian)

	data := gin.H{
		"Secretario":     sec.IDAcademico,
		"Programa":       request.Programanombre,
		"Periodo":        request.PeriodoAcademico,
		"Docentes":       doc,
		"NotasConsejo":   notasConsejo,
		"NotasEvalua":    notasAutoevalua,
		"Cursos":         nombresCursos,
		"IdesCursos":     idecursos,
		"PromedioCursos": promediosCursos,
		"Cantidadest":    estudian,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8081/datos_general", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("error en la petición: %v", resp.Status)
	}

	// Leer el contenido del PDF desde la respuesta
	pdfBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error al leer el cuerpo de la respuesta:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el PDF"})
		return
	}

	config, err := bd.GetConfigSeguridad("copia_controlada")
	if err != nil {
		log.Println("Error al obtener configuración de seguridad:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la configuración"})
		return
	}

	if config.Estado {
		pdfHash := utilities.GeneratePDFHash(pdfBytes)

		// Guardar en la base de datos
		err = bd.SavePDFHash(user, pdfHash)
		if err != nil {
			log.Println("Error al guardar el hash:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el hash"})
			return
		}
	}

	// Enviar el PDF al front-end
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}
