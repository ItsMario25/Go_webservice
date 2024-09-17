package bd

import (
	"log"
	"time"
	"webservice/models"
)

func GetEstudianteid(nombre string) (int, error) {
	db := DBconexion()

	var usuario models.Usuario

	// Consultar el ID del usuario por el nombre
	if err := db.Where("nombre = ?", nombre).First(&usuario).Error; err != nil {
		return 0, err
	}
	var student models.Estudiante
	if err := db.Where("id_user = ?", usuario.IDUser).First(&student).Error; err != nil {
		return 0, err
	}

	return student.CodigoEstudiante, nil
}

func Guardar_evl(evaluacion models.EvaluacionEstudiante) {

	docente := evaluacion.NombreDocente
	estudiante := evaluacion.NombreEvaluador
	curso := evaluacion.NombreCurso
	criterios := evaluacion.Respuestas

	// Aquí podrías guardar las respuestas en la base de datos, por ejemplo.
	log.Printf("Nombre del curso: %s\n", curso)
	log.Printf("Nombre del docente: %s\n", docente)
	log.Printf("Nombre del evaluador: %s\n", estudiante)
	log.Printf("Respuestas: %+v\n", criterios)

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

	cr, err := GetCurso(curso)

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
}
