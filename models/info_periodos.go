package models

type PeriodoEvaluacion struct {
	IDPeriodo string `gorm:"column:id_periodo_evl;primaryKey"`
	Inicio    string `gorm:"column:fecha_inicio"`
	Fin       string `gorm:"column:fecha_final"`
}

func (PeriodoEvaluacion) TableName() string {
	return "periodo_evaluacion"
}

type PeriodoAcademico struct {
	Periodo string `json:"periodo"`
	Inicio  string `json:"inicio"`
	Fin     string `json:"fin"`
}
