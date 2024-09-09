package bd

import (
	"log"
	"webservice/models"
)

func GetCursos_facultad() ([]models.Cursos, error) {
	db := DBconexion()

	var cursos []models.Cursos
	if result := db.Select("id_curso, nombre_curso, id_programa").Find(&cursos); result.Error != nil {
		log.Println(result)
		return nil, result.Error
	} else {
		log.Println(cursos)
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
