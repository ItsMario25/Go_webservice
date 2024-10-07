package bd

import (
	"webservice/models/core"
)

func Get_criterios_bd(tipo string) ([]core.Criterios, error) {
	db := DBconexion()
	var criterio []core.Criterios

	if tipo == "estudiante" {
		formato := "F02"
		if err := db.Select("nombre_criterio").Where("id_formato = ? ", formato).Find(&criterio).Error; err != nil {
			return []core.Criterios{}, err
		} else {
			return criterio, nil
		}
	} else if tipo == "docente" {
		formato := "F03"
		if err := db.Select("nombre_criterio").Where("id_formato = ? ", formato).Find(&criterio).Error; err != nil {
			return []core.Criterios{}, err
		} else {
			return criterio, nil
		}
	} else {
		formato := "F01"
		if err := db.Select("nombre_criterio").Where("id_formato = ? ", formato).Find(&criterio).Error; err != nil {
			return []core.Criterios{}, err
		} else {
			return criterio, nil
		}
	}
}
