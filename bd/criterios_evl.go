package bd

import (
	"webservice/models"
)

func Get_criterios_bd(tipo string) ([]models.Criterios, error) {
	db := DBconexion()
	var criterio []models.Criterios

	if tipo == "estudiante" {
		formato := "F02"
		if err := db.Select("nombre_criterio").Where("id_formato = ? ", formato).Find(&criterio).Error; err != nil {
			return []models.Criterios{}, err
		} else {
			return criterio, nil
		}
	} else if tipo == "docente" {
		formato := "F03"
		if err := db.Select("nombre_criterio").Where("id_formato = ? ", formato).Find(&criterio).Error; err != nil {
			return []models.Criterios{}, err
		} else {
			return criterio, nil
		}
	} else {
		formato := "F01"
		if err := db.Select("nombre_criterio").Where("id_formato = ? ", formato).Find(&criterio).Error; err != nil {
			return []models.Criterios{}, err
		} else {
			return criterio, nil
		}
	}
}
