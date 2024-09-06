package bd

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"webservice/models"
)

func GetPeriod() ([]models.PeriodoEvaluacion, error) {
	db := DBconexion()

	var periodos []models.PeriodoEvaluacion

	if result := db.Select("id_periodo_evl, fecha_inicio, fecha_final").Find(&periodos); result.Error != nil {
		log.Println(result)
		return nil, result.Error
	} else {
		log.Println(periodos)
		return periodos, nil
	}
}

func SetPeriod(periodo string, fechai string, fechaf string) error {
	db := DBconexion()
	numPeriodo, err := strconv.Atoi(periodo)
	if err != nil {
		return err
	}
	year := time.Now().Year()
	idPeriodoAcad := fmt.Sprintf("PA%d-%d", year, numPeriodo)

	periodoAcad := models.PeriodoAcademico{
		IDPeriodoAcad: idPeriodoAcad,
		YearAcad:      year,
		Periodo:       numPeriodo,
	}

	if err := db.Create(&periodoAcad).Error; err != nil {
		return err
	}

	idPeriodoEvl := fmt.Sprintf("PE%d-%d", year, numPeriodo)
	periodoEval := models.PeriodoEvaluacion{
		IDPeriodoEvl:  idPeriodoEvl,
		FechaInicio:   fechai,
		FechaFinal:    fechaf,
		IDPeriodoAcad: idPeriodoAcad,
	}

	if err := db.Create(&periodoEval).Error; err != nil {
		return err
	}

	return nil
}

func GetPeriodoActivo() (*models.PeriodoEvaluacion, error) {
	var periodo models.PeriodoEvaluacion

	today := time.Now()
	db := DBconexion()
	// Buscar el periodo con fecha de inicio menor o igual a la actual y fecha de finalización mayor que la actual
	if err := db.Where("fecha_inicio <= ? AND fecha_final > ?", today, today).First(&periodo).Error; err != nil {
		return nil, err
	} else {
		return &periodo, nil
	}
}

func PutPeriodo(id string, periodo models.PeriodoAc) error {
	db := DBconexion()

	var existingPeriodo models.PeriodoEvaluacion
	if err := db.Where("id_periodo_evl = ?", id).First(&existingPeriodo).Error; err != nil {
		return err
	}

	var periodo_update = models.PeriodoEvaluacion{
		IDPeriodoEvl: periodo.Periodo,
		FechaInicio:  periodo.Inicio,
		FechaFinal:   periodo.Fin,
	}
	// Actualizar el periodo con los nuevos datos
	if err := db.Model(&existingPeriodo).Updates(periodo_update).Error; err != nil {
		return err
	}

	return nil
}
