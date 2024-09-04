package bd

import "webservice/models"

func Get_clave(user string, rol string) (string, error) {
	db := DBconexion()

	var usuario models.Usuario
	var docente models.Docente
	var estudiante models.Estudiante
	var consejo models.ConsejoFacultad
	var secretarioAcademico models.SecretarioAcademico
	var secretarioTecnico models.SecretarioTecnico
	var us string

	if err := db.Where("nombre = ?", user).First(&usuario).Error; err == nil {
		us = usuario.IDUser
	} else {
		return "", err
	}

	if rol == "docente" {
		if err := db.Where("id_user = ?", us).First(&docente).Error; err != nil {
			return docente.ClaveDocente, nil
		} else {
			return "No se encontro usuario asociado", err
		}
	} else if rol == "estudiante" {
		if err := db.Where("id_user = ?", us).First(&estudiante).Error; err != nil {
			return estudiante.ClaveEstudiante, nil
		} else {
			return "No se encontro usuario asociado", err
		}
	} else if rol == "consejo_facultad" {
		if err := db.Where("id_user = ?", us).First(&consejo).Error; err != nil {
			return consejo.ClaveConsejo, nil
		} else {
			return "No se encontro usuario asociado", err
		}
	} else if rol == "secretario_academico" {
		if err := db.Where("id_user = ?", us).First(&secretarioAcademico).Error; err != nil {
			return secretarioAcademico.ClaveAcademico, nil
		} else {
			return "No se encontro usuario asociado", err
		}
	} else {
		if err := db.Where("id_user = ?", us).First(&secretarioTecnico).Error; err != nil {
			return secretarioTecnico.ClaveSecretario, nil
		} else {
			return "No se encontro usuario asociado", err
		}
	}
}
