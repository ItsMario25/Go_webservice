package bd

import (
	"fmt"
	"log"
	"sort"
	"time"
	"webservice/models"
)

func Guardar_evl(evaluacion models.FormatoEvaluacion) {
	db := DBconexion()

	docente := evaluacion.NombreDocente
	estudiante := evaluacion.NombreEvaluador
	curso := evaluacion.NombreCurso
	criterios := evaluacion.Respuestas

	dc, err := GetDocente(docente)
	if err != nil {
		log.Println(err)
	}

	est, err := GetUsuarioid(estudiante)
	if err != nil {
		log.Println(err)
	}

	cr, err := GetCursoxnombre(curso)
	if err != nil {
		log.Println(err)
	}

	fechaActual := time.Now().Format("2006-01-02")

	periodoac, err := GetPeriodoActivo()
	if err != nil {
		log.Panic("Periodo de evaluacion no encontrado")
	}
	periodoacc, err := GetPeriodoAcActivo()
	if err != nil {
		log.Panic("Periodo de academico no encontrado")
	}

	ejer, err := GetEjerciendo(dc.IDDocente, cr.IDCurso, periodoacc.IDPeriodoAcad)
	if err != nil {
		log.Println(err)
	}

	var claves []int
	for clave := range criterios {
		claves = append(claves, clave)
	}
	sort.Ints(claves)

	for _, clave := range claves {
		valor := criterios[clave]
		clave = clave + 1
		if clave < 10 {
			criterio := fmt.Sprintf("C0%d", clave)

			nuevaEvaluacion := models.Evaluacion{
				FechaDiligenciada: fechaActual,
				Calificacion:      valor,
				IdCriterio:        criterio,
				IdUser:            est,
				IdPeriodoEvl:      periodoac.IDPeriodoEvl,
				IdEjerciendo:      ejer,
			}
			if err := db.Create(&nuevaEvaluacion).Error; err != nil {
				fmt.Println("Error al insertar la evaluación:", err)
			} else {
				fmt.Println("Evaluación insertada con éxito:", nuevaEvaluacion.IdEvaluacion)
			}
		} else {
			criterio := fmt.Sprintf("C%d", clave)
			nuevaEvaluacion := models.Evaluacion{
				FechaDiligenciada: fechaActual,
				Calificacion:      valor,
				IdCriterio:        criterio,
				IdUser:            est,
				IdPeriodoEvl:      periodoac.IDPeriodoEvl,
				IdEjerciendo:      ejer,
			}
			if err := db.Create(&nuevaEvaluacion).Error; err != nil {
				fmt.Println("Error al insertar la evaluación:", err)
			} else {
				fmt.Println("Evaluación insertada con éxito:", nuevaEvaluacion.IdEvaluacion)
			}
		}
	}
}

func Guardar_evl_docente(autoevaluacion models.EvaluacionDocente) {
	db := DBconexion()

	curso := autoevaluacion.NombreCurso
	docente := autoevaluacion.NombreDocente
	criterios := autoevaluacion.Respuestas

	dc, err := GetDocente(docente)
	if err != nil {
		log.Println(err)
	}

	evl, err := GetUsuarioid(docente)
	if err != nil {
		log.Println(err)
	}

	cr, err := GetCursoxnombre(curso)
	if err != nil {
		log.Println(err)
	}

	fechaActual := time.Now().Format("2006-01-02")

	periodoac, err := GetPeriodoActivo()
	if err != nil {
		log.Panic("Periodo de evaluacion no encontrado")
	}
	periodoacc, err := GetPeriodoAcActivo()
	if err != nil {
		log.Panic("Periodo de academico no encontrado")
	}

	ejer, err := GetEjerciendo(dc.IDDocente, cr.IDCurso, periodoacc.IDPeriodoAcad)
	if err != nil {
		log.Println(err)
	}

	var claves []int
	for clave := range criterios {
		claves = append(claves, clave)
	}
	sort.Ints(claves)

	for _, clave := range claves {
		valor := criterios[clave]
		clave = clave + 47
		criterio := fmt.Sprintf("C%d", clave)

		log.Println(criterio)
		nuevaEvaluacion := models.Evaluacion{
			FechaDiligenciada: fechaActual,
			Calificacion:      valor,
			IdCriterio:        criterio,
			IdUser:            evl,
			IdPeriodoEvl:      periodoac.IDPeriodoEvl,
			IdEjerciendo:      ejer,
		}

		if err := db.Create(&nuevaEvaluacion).Error; err != nil {
			fmt.Println("Error al insertar la evaluación:", err)
		} else {
			fmt.Println("Evaluación insertada con éxito:", nuevaEvaluacion.IdEvaluacion)
		}

	}
}

func Guardar_evl_facultad(evaluacion models.FormatoEvaluacion) {
	db := DBconexion()

	docente := evaluacion.NombreDocente
	consejo := evaluacion.NombreEvaluador
	curso := evaluacion.NombreCurso
	criterios := evaluacion.Respuestas

	dc, err := GetDocente(docente)
	if err != nil {
		log.Println(err)
	}

	cr, err := GetCursoxnombre(curso)
	if err != nil {
		log.Println(err)
	}

	periodoacc, err := GetPeriodoAcActivo()
	if err != nil {
		log.Panic("Periodo de academico no encontrado")
	}

	ejer, err := GetEjerciendo(dc.IDDocente, cr.IDCurso, periodoacc.IDPeriodoAcad)
	if err != nil {
		log.Println(err)
	}

	periodoac, err := GetPeriodoActivo()
	if err != nil {
		log.Panic("Periodo de evaluacion no encontrado")
	}

	cons, err := GetUsuarioid(consejo)
	if err != nil {
		log.Println(err)
	}

	var claves []int
	for clave := range criterios {
		claves = append(claves, clave)
	}
	sort.Ints(claves)

	fechaActual := time.Now().Format("2006-01-02")

	for _, clave := range claves {
		valor := criterios[clave]
		clave = clave + 24
		criterio := fmt.Sprintf("C%d", clave)

		log.Println(criterio)
		nuevaEvaluacion := models.Evaluacion{
			FechaDiligenciada: fechaActual,
			Calificacion:      valor,
			IdCriterio:        criterio,
			IdUser:            cons,
			IdPeriodoEvl:      periodoac.IDPeriodoEvl,
			IdEjerciendo:      ejer,
		}

		if err := db.Create(&nuevaEvaluacion).Error; err != nil {
			log.Println("Error al insertar la evaluación:", err)
		} else {
			log.Println("Evaluación insertada con éxito:", nuevaEvaluacion.IdEvaluacion)
		}

	}
}

func Get_materias_evaluadas(evaluador string, periodo string) {
	db := DBconexion()
	var evaluado []models.Evaluacion

	iduser, err := GetUsuarioid(evaluador)

	if err != nil {
		log.Println("ERROR DE USUARIO")
	}

	if err := db.Where("id_user = ? and id_periodo_evl = ? ", iduser, periodo).Find(&evaluado).Error; err != nil {
		log.Println("error de lectura")
	}
}
