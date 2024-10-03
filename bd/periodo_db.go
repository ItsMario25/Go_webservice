package bd

import (
	"fmt"
	"strconv"
	"time"
	"webservice/models"
)

func GetPeriod() ([]models.PeriodoEvaluacion, error) {
	db := DBconexion()

	var periodos []models.PeriodoEvaluacion

	if result := db.Select("id_periodo_evl, fecha_inicio, fecha_final").Find(&periodos); result.Error != nil {
		return nil, result.Error
	} else {
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

func GetPeriodoAcActivo() (*models.PeriodoAcademico, error) {
	var periodoAc models.PeriodoAcademico

	today := time.Now()
	db := DBconexion()
	// Buscar el periodo ACADEMICO con fecha de inicio menor o igual a la actual y fecha de finalización mayor que la actual
	if err := db.Where("fecha_inicial <= ? AND fecha_final > ?", today, today).First(&periodoAc).Error; err != nil {
		return nil, err
	} else {
		return &periodoAc, nil
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

func GetPeriodo_facultad(id int) ([]models.PeriodoEvaluacion, error) {
	db := DBconexion()
	var periodos []models.PeriodoEvaluacion

	err := db.Table("periodo_evaluacion").
		Select("DISTINCT REPLACE(periodo_evaluacion.id_periodo_evl, 'PE', '') as id_periodo_evl, periodo_evaluacion.fecha_inicio, periodo_evaluacion.fecha_final").
		Joins("JOIN periodo_academico ON periodo_evaluacion.id_periodo_acad = periodo_academico.id_periodo_acad").
		Joins("JOIN ejerciendo ON periodo_academico.id_periodo_acad = ejerciendo.id_periodo_acad").
		Joins("JOIN curso ON ejerciendo.id_curso = curso.id_curso").
		Joins("JOIN programa ON curso.id_programa = programa.id_programa").
		Where("programa.id_facultad = ? AND periodo_evaluacion.fecha_final <= NOW()", id).
		Scan(&periodos).Error

	if err != nil {
		return nil, err
	}

	return periodos, nil
}
