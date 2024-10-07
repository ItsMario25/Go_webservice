package bd

import (
	"fmt"
	"log"
	"webservice/models/core"
)

func GetDocentes() ([]core.DocenteConUsuario, error) {
	db := DBconexion()

	var docentes []core.DocenteConUsuario

	// Realizar la consulta con un JOIN
	err := db.Table("docente").
		Select("docente.id_docente, usuarios.nombre").
		Joins("inner join usuarios on usuarios.id_user = docente.id_user").
		Scan(&docentes).Error

	if err != nil {
		log.Fatalf("Error al obtener los docentes con usuario: %v", err)
		return []core.DocenteConUsuario{}, err
	} else {
		// Imprimir los resultados
		for _, docente := range docentes {
			fmt.Printf("ID Docente: %d, Nombre Usuario: %s\n", docente.IdDocente, docente.Nombre)
		}
		return docentes, nil
	}
}

func GetDocente(nombre string) (core.Docente, error) {
	db := DBconexion()

	var docente core.Docente

	err := db.Table("docente").
		Select("docente.id_docente").
		Joins("inner join usuarios on usuarios.id_user = docente.id_user").
		Where("usuarios.nombre = ?", nombre).
		Scan(&docente).Error

	if err != nil {
		log.Fatalf("Error al obtener los docentes con usuario: %v", err)
		return core.Docente{}, err
	} else {
		return docente, nil
	}
}

func GetDocentesActuales(codigoEstudiante int) ([]core.DocenteCurso, error) {
	db := DBconexion()

	periodoVigente, err := GetPeriodoAcActivo()
	if err != nil {
		return nil, err
	}
	// Variable para almacenar los cursos que está cursando el estudiante en el periodo vigente
	var cursosEstudiante []int

	// Obtener los cursos que está cursando el estudiante en el periodo vigente
	if err := db.Model(&core.Cursando{}).
		Where("codigo_estudiante = ? AND id_periodo_acad = ?", codigoEstudiante, periodoVigente.IDPeriodoAcad).
		Pluck("id_curso", &cursosEstudiante).Error; err != nil {
		return nil, err
	}

	// Si no hay cursos, retornar un error o array vacío
	if len(cursosEstudiante) == 0 {
		return []core.DocenteCurso{}, nil
	}

	// Variable para almacenar los ID de los docentes que están ejerciendo esos mismos cursos
	var docentes []int

	// Obtener los docentes que están ejerciendo esos mismos cursos en el periodo vigente
	if err := db.Model(&core.Ejerciendo{}).
		Where("id_curso IN (?) AND id_periodo_acad = ?", cursosEstudiante, periodoVigente.IDPeriodoAcad).
		Pluck("id_docente", &docentes).Error; err != nil {
		return nil, err
	}

	// Obtener la información de los docentes
	var docentesInfo []core.Docente
	if err := db.Where("id_docente IN (?)", docentes).Find(&docentesInfo).Error; err != nil {
		return nil, err
	}

	// Obtener la información de los cursos
	var cursosInfo []core.Cursos
	if err := db.Where("id_curso IN (?)", cursosEstudiante).Find(&cursosInfo).Error; err != nil {
		return nil, err
	}

	// Crear un mapa de docentes con ID y nombre
	docentesMap := make(map[int]string)
	for _, docente := range docentesInfo {
		var usuario core.Usuario
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
	var resultado []core.DocenteCurso
	for _, docente := range docentesInfo {
		for _, cursoID := range cursosEstudiante {
			var ejerciendo core.Ejerciendo
			if err := db.Where("id_docente = ? AND id_curso = ? AND id_periodo_acad = ?", docente.IDDocente, cursoID, periodoVigente.IDPeriodoAcad).First(&ejerciendo).Error; err == nil {
				resultado = append(resultado, core.DocenteCurso{
					NombreDocente: docentesMap[docente.IDDocente],
					NombreCurso:   cursosMap[cursoID],
				})
			}
		}
	}

	return resultado, nil
}

func GetDocentesYCursosByEjerciendo(ejerciendo []core.Ejerciendo) ([]core.DocenteCurso, error) {
	var docentesCursos []core.DocenteCurso
	db := DBconexion()

	for _, ejerce := range ejerciendo {
		var docente core.Docente
		var curso core.Cursos
		var usuario core.Usuario

		// Obtener el curso
		if err := db.Where("id_curso = ?", ejerce.IDcurso).First(&curso).Error; err != nil {
			return nil, err
		}

		// Obtener el docente
		if err := db.Where("id_docente = ?", ejerce.IDdocente).First(&docente).Error; err != nil {
			return nil, err
		}

		// Obtener el nombre del docente desde la tabla Usuarios
		if err := db.Where("id_user = ?", docente.IDUser).First(&usuario).Error; err != nil {
			return nil, err
		}

		// Llenar la información del modelo DocenteCurso
		docentesCursos = append(docentesCursos, core.DocenteCurso{
			NombreDocente: usuario.Nombre,
			NombreCurso:   curso.NombreCurso,
		})
	}

	return docentesCursos, nil
}
