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
