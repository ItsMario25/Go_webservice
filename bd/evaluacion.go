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

	fechaActual := time.Now().Format("2006-01-02")

	periodoac, err := GetPeriodoActivo()
	if err != nil {
		log.Panic("Periodo de evaluacion no encontrado")
	}
	periodoacc, err := GetPeriodoAcActivo()
	if err != nil {
		log.Panic("Periodo de academico no encontrado")
	}

	ejer, err := GetEjerciendoporDocente(dc.IDDocente, periodoacc.IDPeriodoAcad)
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

func Guardar_evl_facultad(evaluacion models.FormatoEvaluacionFacultad) {
	db := DBconexion()

	docente := evaluacion.NombreDocente
	consejo := evaluacion.NombreEvaluador
	criterios := evaluacion.Respuestas

	dc, err := GetDocente(docente)
	if err != nil {
		log.Println(err)
	}

	periodoacc, err := GetPeriodoAcActivo()
	if err != nil {
		log.Panic("Periodo de academico no encontrado")
	}

	ejer, err := GetEjerciendoporDocente(dc.IDDocente, periodoacc.IDPeriodoAcad)
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

func Get_materias_evaluadas(evaluador string, periodo string) ([]string, error) {
	db := DBconexion()
	var evaluado []int
	var idCursos []int
	var nombresCursos []string

	iduser, err := GetUsuarioid(evaluador)

	if err != nil {
		log.Println("ERROR DE USUARIO")
	}

	if err := db.Model(&models.Evaluacion{}).Select("DISTINCT id_ejerciendo").
		Where("id_user = ? AND id_periodo_evl = ?", iduser, periodo).Pluck("id_ejerciendo", &evaluado).Error; err != nil {

		log.Println("Error de lectura:", err)
		return nil, err
	}

	if err := db.Model(&models.Ejerciendo{}).Select("DISTINCT id_curso").
		Where("id_ejerciendo IN ?", evaluado).Pluck("id_curso", &idCursos).Error; err != nil {

		log.Println("Error al obtener los IdCurso:", err)
		return nil, err
	}

	if err := db.Model(&models.Cursos{}).Select("nombre_curso").
		Where("id_curso IN ?", idCursos).Pluck("nombre_curso", &nombresCursos).Error; err != nil {

		log.Println("Error al obtener los nombres de los cursos:", err)
		return nil, err
	}

	return nombresCursos, nil
}

func Get_docentes_evl(evaluador string, periodo string) ([]string, error) {
	db := DBconexion()
	var evaluado []int
	var idDocente []int
	var nombresDocente []string
	var nombres []string

	iduser, err := GetUsuarioid(evaluador)

	if err != nil {
		log.Println("ERROR DE USUARIO")
	}

	if err := db.Model(&models.Evaluacion{}).Select("DISTINCT id_ejerciendo").
		Where("id_user = ? AND id_periodo_evl = ?", iduser, periodo).Pluck("id_ejerciendo", &evaluado).Error; err != nil {

		log.Println("Error de lectura:", err)
		return nil, err
	}

	if err := db.Model(&models.Ejerciendo{}).Select("DISTINCT id_docente").
		Where("id_ejerciendo IN ?", evaluado).Pluck("id_docente", &idDocente).Error; err != nil {

		log.Println("Error al obtener los id de docentes:", err)
		return nil, err
	}

	if err := db.Model(&models.Docente{}).Select("id_user").
		Where("id_docente IN ?", idDocente).Pluck("id_user", &nombresDocente).Error; err != nil {

		log.Println("Error al obtener los id de usuarios:", err)
		return nil, err
	}

	if err := db.Model(&models.Usuario{}).Select("nombre").
		Where("id_user IN ?", nombresDocente).Pluck("nombre", &nombres).Error; err != nil {

		log.Println("Error al obtener los id de usuarios:", err)
		return nil, err
	}

	return nombres, nil
}
