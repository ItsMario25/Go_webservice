package bd

import (
	"webservice/models"
)

func GetCursos_facultad() ([]models.Cursos, error) {
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
