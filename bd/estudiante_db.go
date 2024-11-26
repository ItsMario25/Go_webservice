package bd

import (
	"webservice/models/core"
)

func GetEstudianteid(nombre string) (int, error) {
	db := DBconexion()

	var usuario core.Usuario

	// Consultar el ID del usuario por el nombre
	if err := db.Where("nombre = ?", nombre).First(&usuario).Error; err != nil {
		return 0, err
	}
	var student core.Estudiante
	if err := db.Where("id_user = ?", usuario.IDUser).First(&student).Error; err != nil {
		return 0, err
	}

	return student.CodigoEstudiante, nil
}
