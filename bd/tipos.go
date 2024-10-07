package bd

import (
	"log"
	"webservice/models/core"
)

func GetTipos() ([]core.Tipos, error) {
	db := DBconexion()

	var tipos []core.Tipos
	err := db.Table("tipo").
		Select("id_tipo, nombre_tipo").
		Scan(&tipos).Error

	if err != nil {
		log.Fatalf("Error al obtener los docentes con usuario: %v", err)
		return []core.Tipos{}, err
	} else {
		return tipos, nil
	}
}
