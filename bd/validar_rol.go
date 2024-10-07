package bd

import (
	"log"
	"time"
	"webservice/models/core"
)

func Get_clave(user string, rol string) (string, error) {
	db := DBconexion()

	var usuario core.Usuario
	var docente core.Docente
	var estudiante core.Estudiante
	var consejo core.ConsejoFacultad
	var secretarioAcademico core.SecretarioAcademico
	var secretarioTecnico core.SecretarioTecnico
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

func StoreTokenInDB(correo string, token string) error {
	db := DBconexion()
	creadoEn := time.Now()
	expiracion := creadoEn.Add(15 * time.Minute)

	err := db.Exec(`
        INSERT INTO token_verificacion (token, creado_en, expiracion, correo) 
        VALUES (?, ?, ?, ?)`, token, creadoEn, expiracion, correo).Error
	if err != nil {
		return err
	}
	return nil
}

func Verificar_token_correo(token string) error {
	db := DBconexion()
	var tokenVerificado struct {
		Token      string
		CreadoEn   time.Time
		Expiracion time.Time // Ahora guardamos la fecha de expiración directamente
	}

	err := db.Raw(`
        SELECT token, creado_en, expiracion 
        FROM token_verificacion 
        WHERE token = ? `, token).Scan(&tokenVerificado).Error

	log.Println(tokenVerificado)
	if err != nil {
		log.Println("Error escaneando los valores:", err)
		return err
	}
	// Token es válido
	log.Println("Token verificado correctamente:", tokenVerificado)
	return nil
}
