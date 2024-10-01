package bd

import (
	"log"
	"webservice/models"
)

func GetTipos() ([]models.Tipos, error) {
	db := DBconexion()

	var tipos []models.Tipos
	err := db.Table("tipo").
		Select("id_tipo, nombre_tipo").
		Scan(&tipos).Error

	if err != nil {
		log.Fatalf("Error al obtener los docentes con usuario: %v", err)
		return []models.Tipos{}, err
	} else {
		return tipos, nil
	}
}
