package bd

import (
	"fmt"
	"log"
	"webservice/models"
)

func GetDocentes() ([]models.DocenteConUsuario, error) {
	db := DBconexion()

	var docentes []models.DocenteConUsuario

	// Realizar la consulta con un JOIN
	err := db.Table("docente").
		Select("docente.id_docente, usuarios.nombre").
		Joins("inner join usuarios on usuarios.id_user = docente.id_user").
		Scan(&docentes).Error

	if err != nil {
		log.Fatalf("Error al obtener los docentes con usuario: %v", err)
		return []models.DocenteConUsuario{}, err
	} else {
		// Imprimir los resultados
		for _, docente := range docentes {
			fmt.Printf("ID Docente: %d, Nombre Usuario: %s\n", docente.IdDocente, docente.Nombre)
		}
		return docentes, nil
	}
}

func GetDocente(nombre string) (models.Docente, error) {
	db := DBconexion()

	var docente models.Docente

	err := db.Table("docente").
		Select("docente.id_docente").
		Joins("inner join usuarios on usuarios.id_user = docente.id_user").
		Where("usuarios.nombre = ?", nombre).
		Scan(&docente).Error

	if err != nil {
		log.Fatalf("Error al obtener los docentes con usuario: %v", err)
		return models.Docente{}, err
	} else {
		return docente, nil
	}
}

func GetDocentesActuales(codigoEstudiante int) ([]models.DocenteCurso, error) {
	db := DBconexion()

	periodoVigente, err := GetPeriodoAcActivo()
	if err != nil {
		return nil, err
	}
	// Variable para almacenar los cursos que está cursando el estudiante en el periodo vigente
	var cursosEstudiante []int

	// Obtener los cursos que está cursando el estudiante en el periodo vigente
	if err := db.Model(&models.Cursando{}).
		Where("codigo_estudiante = ? AND id_periodo_acad = ?", codigoEstudiante, periodoVigente.IDPeriodoAcad).
		Pluck("id_curso", &cursosEstudiante).Error; err != nil {
		return nil, err
	}

	// Si no hay cursos, retornar un error o array vacío
	if len(cursosEstudiante) == 0 {
		return []models.DocenteCurso{}, nil
	}

	// Variable para almacenar los ID de los docentes que están ejerciendo esos mismos cursos
	var docentes []int

	// Obtener los docentes que están ejerciendo esos mismos cursos en el periodo vigente
	if err := db.Model(&models.Ejerciendo{}).
		Where("id_curso IN (?) AND id_periodo_acad = ?", cursosEstudiante, periodoVigente.IDPeriodoAcad).
		Pluck("id_docente", &docentes).Error; err != nil {
		return nil, err
	}

	// Obtener la información de los docentes
	var docentesInfo []models.Docente
	if err := db.Where("id_docente IN (?)", docentes).Find(&docentesInfo).Error; err != nil {
		return nil, err
	}

	// Obtener la información de los cursos
	var cursosInfo []models.Cursos
	if err := db.Where("id_curso IN (?)", cursosEstudiante).Find(&cursosInfo).Error; err != nil {
		return nil, err
	}

	// Crear un mapa de docentes con ID y nombre
	docentesMap := make(map[int]string)
	for _, docente := range docentesInfo {
		var usuario models.Usuario
		if err := db.Where("id_user = ?", docente.IDUser).First(&usuario).Error; err != nil {
			return nil, err
		}
		docentesMap[docente.IDDocente] = usuario.Nombre
	}

	// Crear un mapa de cursos con ID y nombre
	cursosMap := make(map[int]string)
	for _, curso := range cursosInfo {
		cursosMap[curso.IDCurso] = curso.NombreCurso
	}

	// Combinar la información
	var resultado []models.DocenteCurso
	for _, docente := range docentesInfo {
		for _, cursoID := range cursosEstudiante {
			var ejerciendo models.Ejerciendo
			if err := db.Where("id_docente = ? AND id_curso = ? AND id_periodo_acad = ?", docente.IDDocente, cursoID, periodoVigente.IDPeriodoAcad).First(&ejerciendo).Error; err == nil {
				resultado = append(resultado, models.DocenteCurso{
					NombreDocente: docentesMap[docente.IDDocente],
					NombreCurso:   cursosMap[cursoID],
				})
			}
		}
	}

	return resultado, nil
}
