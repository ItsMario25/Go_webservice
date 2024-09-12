package bd

import "webservice/models"

func GetEstudianteid(nombre string) (int, error) {
	db := DBconexion()

	var usuario models.Usuario

	// Consultar el ID del usuario por el nombre
	if err := db.Where("nombre = ?", nombre).First(&usuario).Error; err != nil {
		return 0, err
	}
	var student models.Estudiante
	if err := db.Where("id_user = ?", usuario.IDUser).First(&student).Error; err != nil {
		return 0, err
	}

	return student.CodigoEstudiante, nil
}
