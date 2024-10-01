package bd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"webservice/models"
	"webservice/utilidades"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
	}

	DB = db
	return DB
}

func DBconexion() *gorm.DB {
	sqlDB, err := DB.DB() // Obtén la conexión subyacente de *sql.DB
	if err != nil {
		log.Printf("Error al obtener la conexión a la base de datos: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Printf("Conexión a la base de datos no válida: %v", err)
	}

	return DB
}

func ValidarUsuario(idUser string, clave string) (bool, string, string, error) {
	db := DBconexion()

	var usuario models.Usuario
	var rol string
	var us string
	var claveValida bool
	var docente models.Docente
	var estudiante models.Estudiante
	var consejo models.ConsejoFacultad
	var secretarioAcademico models.SecretarioAcademico
	var secretarioTecnico models.SecretarioTecnico

	if err := db.Where("id_docente = ?", idUser).First(&docente).Error; err == nil {
		if utilidades.ValidarPassword(clave, docente.ClaveDocente) {
			rol = "docente"
			us = docente.IDUser
			claveValida = true
		}
	} else if err := db.Where("codigo_estudiante = ?", idUser).First(&estudiante).Error; err == nil {
		if utilidades.ValidarPassword(clave, estudiante.ClaveEstudiante) {
			rol = "estudiante"
			us = estudiante.IDUser
			claveValida = true
		}
	} else if err := db.Where("id_consejo = ?", idUser).First(&consejo).Error; err == nil {
		if utilidades.ValidarPassword(clave, consejo.ClaveConsejo) {
			rol = "consejo_facultad"
			us = consejo.IDUser
			claveValida = true
		}
	} else if err := db.Where("id_academico = ?", idUser).First(&secretarioAcademico).Error; err == nil {
		if utilidades.ValidarPassword(clave, secretarioAcademico.ClaveAcademico) {
			rol = "secretario_academico"
			us = secretarioAcademico.IDUser
			claveValida = true
			/*token := utilidades.GenerateToken()
			if err := utilidades.SendTokenEmail(secretarioAcademico.Correo, token); err != nil {
				log.Println(err)
			}*/
		}
	} else if err := db.Where("id_secret = ?", idUser).First(&secretarioTecnico).Error; err == nil {
		if utilidades.ValidarPassword(clave, secretarioTecnico.ClaveSecretario) {
			rol = "secretario_tecnico"
			us = secretarioTecnico.IDUser
			claveValida = true
			/*token := utilidades.GenerateToken()
			if err := utilidades.SendTokenEmail(secretarioTecnico.Correo, token); err != nil {
				log.Println(err)
			}*/
		}
	}

	if !claveValida {
		return false, "", "", errors.New("usuario o clave incorrectos")
	}

	if err := db.Where("id_user = ?", us).First(&usuario).Error; err != nil {
		return false, "", "", errors.New("no se encontró el usuario")
	}

	return true, usuario.Nombre, rol, nil
}
