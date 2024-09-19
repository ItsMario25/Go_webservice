package bd

import (
	"log"
	"strconv"
	"webservice/models"
)

func GetEjerciendo(docente int, curso int, periodo string) (int, error) {
	db := DBconexion()

	var ejerciendoID int

	err := db.Table("ejerciendo").
		Select("id_ejerciendo").
		Where("id_docente = ? AND id_curso = ? AND id_periodo_acad = ?", docente, curso, periodo).
		Scan(&ejerciendoID).Error

	if err != nil {
		return 0, err
	} else {
		return ejerciendoID, nil
	}

}

func GetEjerciendoActual(docente int) ([]models.Cursos, error) {
	db := DBconexion()

	var cursosId []int

	periodo, err := GetPeriodoActivo()
	if err != nil {
		log.Println(err)
	}

	err = db.Table("ejerciendo").
		Select("id_curso").
		Where("id_docente = ? AND id_periodo_acad = ?", docente, periodo.IDPeriodoAcad).
		Scan(&cursosId).Error

	if err != nil {
		log.Println(err)
	}

	var curse []models.Cursos

	for _, id := range cursosId {
		idCurso := strconv.Itoa(id)

		curso, err := GetCurso(idCurso)
		if err != nil {
			return []models.Cursos{}, err
		}
		curse = append(curse, curso)
	}

	return curse, nil
}

func SetEjerciendo(docente int, curso int) error {
	db := DBconexion()
	periodo, err := GetPeriodoAcActivo()

	if err != nil {
		return err
	}

	var setejer = models.Ejerciendo{
		IDdocente: docente,
		IDcurso:   curso,
		IDperiodo: periodo.IDPeriodoAcad,
	}

	if err := db.Create(&setejer).Error; err != nil {
		return err
	}

	return nil
}
