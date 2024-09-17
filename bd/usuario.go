package bd

import "webservice/models"

func GetUsuarioid(nombre string) (string, error) {
	db := DBconexion()

	var usuario models.Usuario

	// Consultar el ID del usuario por el nombre
	if err := db.Where("nombre = ?", nombre).First(&usuario).Error; err != nil {
		return "", err
	}

	return usuario.IDUser, nil
}
