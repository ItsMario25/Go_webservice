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

func Get_Secretario(idus string) (models.SecretarioAcademico, error) {
	db := DBconexion()

	var usuario models.SecretarioAcademico

	if err := db.Where("id_user = ?", idus).First(&usuario).Error; err != nil {
		return models.SecretarioAcademico{}, err
	}

	return usuario, nil
}
