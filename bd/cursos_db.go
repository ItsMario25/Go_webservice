package bd

import (
	"webservice/models"
)

func GetCursos() ([]models.Cursos, error) {
	db := DBconexion()

	var cursos []models.Cursos
	if result := db.Select("id_curso, nombre_curso, id_programa").Find(&cursos); result.Error != nil {
		return nil, result.Error
	} else {
		return cursos, nil
	}
}

func GetCurso(idCurso string) (models.Cursos, error) {
	db := DBconexion()

	var cursos models.Cursos

	if err := db.Where("id_curso = ? ", idCurso).First(&cursos).Error; err != nil {
		return models.Cursos{}, err
	} else {
		return cursos, nil
	}
}

func GetCursosAsignados() ([]int, error) {
	db := DBconexion()

	var cursosAsignados []int
	periodoActual, err := GetPeriodoAcActivo()

	if err != nil {
		return nil, err
	}

	if err := db.Model(&models.Ejerciendo{}).Where("id_periodo_acad = ?", periodoActual.IDPeriodoAcad).Pluck("id_curso", &cursosAsignados).Error; err != nil {
		return nil, err
	}

	return cursosAsignados, nil
}

func GetCursoxnombre(nomCurso string) (models.Cursos, error) {
	db := DBconexion()

	var cursos models.Cursos

	if err := db.Where("nombre_curso = ? ", nomCurso).First(&cursos).Error; err != nil {
		return models.Cursos{}, err
	} else {
		return cursos, nil
	}
}

func GetCursosByPeriodo(idPeriodoAcad string) ([]models.Ejerciendo, error) {
	var ejerciendo []models.Ejerciendo
	db := DBconexion()

	// Obtener todos los cursos del periodo académico activo
	if err := db.Where("id_periodo_acad = ?", idPeriodoAcad).Find(&ejerciendo).Error; err != nil {
		return nil, err
	}
	return ejerciendo, nil
}

func GetCursosPorFacultad(idFacultad int) ([]models.Cursos, error) {
	var cursos []models.Cursos
	db := DBconexion()

	// Obtener cursos que pertenecen a la facultad
	if err := db.Joins("left join programa ON curso.id_programa = programa.id_programa").
		Where("programa.id_facultad = ?", idFacultad).Find(&cursos).Error; err != nil {
		return nil, err
	}
	return cursos, nil
}
