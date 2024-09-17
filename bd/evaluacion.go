package bd

import (
	"fmt"
	"log"
	"time"
	"webservice/models"
)

func Guardar_evl(evaluacion models.EvaluacionEstudiante) {
	db := DBconexion()

	docente := evaluacion.NombreDocente
	estudiante := evaluacion.NombreEvaluador
	curso := evaluacion.NombreCurso
	criterios := evaluacion.Respuestas

	dc, err := GetDocente(docente)

	if err != nil {
		log.Println(err)
	}

	log.Println(dc.IDDocente)

	est, err := GetUsuarioid(estudiante)

	if err != nil {
		log.Println(err)
	}

	log.Println(est)

	cr, err := GetCursoxnombre(curso)

	if err != nil {
		log.Println(err)
	}

	log.Println(cr.IDCurso)

	now := time.Now()
	fechaActual := now.Format("2006-01-02")
	log.Println("Fecha actual:", fechaActual)

	periodoac, err := GetPeriodoActivo()
	if err != nil {
		log.Panic("Periodo de evaluacion no encontrado")
	}

	periodoacc, err := GetPeriodoAcActivo()
	if err != nil {
		log.Panic("Periodo de evaluacion no encontrado")
	}

	log.Println(periodoacc.IDPeriodoAcad)
	log.Println(periodoac.IDPeriodoEvl)

	ejer, err := GetEjerciendo(dc.IDDocente, cr.IDCurso, periodoacc.IDPeriodoAcad)

	if err != nil {
		log.Println(err)
	}

	log.Println(ejer)

	for clave, valor := range criterios {
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
