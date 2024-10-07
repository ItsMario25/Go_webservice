package bd

import "webservice/models/core"

func GetUsuarioid(nombre string) (string, error) {
	db := DBconexion()

	var usuario core.Usuario

	// Consultar el ID del usuario por el nombre
	if err := db.Where("nombre = ?", nombre).First(&usuario).Error; err != nil {
		return "", err
	}

	return usuario.IDUser, nil
}

func Get_Secretario(idus string) (core.SecretarioAcademico, error) {
	db := DBconexion()

	var usuario core.SecretarioAcademico

	if err := db.Where("id_user = ?", idus).First(&usuario).Error; err != nil {
		return core.SecretarioAcademico{}, err
	}

	return usuario, nil
}
