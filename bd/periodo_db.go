package bd

import (
	"log"
	"webservice/models"
)

func GetPeriod() ([]models.PeriodoEvaluacion, error) {
	db := DBconexion()

	var periodos []models.PeriodoEvaluacion

	if result := db.Find(&periodos); result.Error != nil {
		log.Println(result)
		return nil, result.Error
	} else {
		log.Println(periodos)
		return periodos, nil
	}
}
