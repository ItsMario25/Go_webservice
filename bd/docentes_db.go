package bd

import (
	"fmt"
	"log"
	"webservice/models"
)

func GetDocentes() ([]models.DocenteConUsuario, error) {
	db := DBconexion()

	var docentes []models.DocenteConUsuario

	// Realizar la consulta con un JOIN
	err := db.Table("docente").
		Select("docente.id_docente, usuarios.nombre").
		Joins("inner join usuarios on usuarios.id_user = docente.id_user").
		Scan(&docentes).Error

	if err != nil {
		log.Fatalf("Error al obtener los docentes con usuario: %v", err)
		return []models.DocenteConUsuario{}, err
	} else {
		// Imprimir los resultados
		for _, docente := range docentes {
			fmt.Printf("ID Docente: %d, Nombre Usuario: %s\n", docente.IdDocente, docente.Nombre)
		}
		return docentes, nil
	}

}

func SetEjerciendo(docente int, curso int) error {
	db := DBconexion()
	periodo, err := GetPeriodoAcActivo()

	if err != nil {
		log.Println("Cagaste")
	}

	var setejer = models.Ejerciendo{
		IDdocente: docente,
		IDcurso:   curso,
		IDperiodo: periodo.IDPeriodoAcad,
	}

	if err := db.Create(&setejer).Error; err != nil {
		return err
	}

	return nil
}
