package bd

import "webservice/models"

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
