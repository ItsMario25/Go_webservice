package models

type PeriodoEvaluacion struct {
	IDPeriodoEvl  string `gorm:"primaryKey" json:"id_periodo_evl"`
	FechaInicio   string `json:"fecha_inicio"`
	FechaFinal    string `json:"fecha_final"`
	IDPeriodoAcad string `json:"id_periodo_acad"`
}

func (PeriodoEvaluacion) TableName() string {
	return "periodo_evaluacion"
}

type PeriodoAc struct {
	Periodo string `json:"periodo"`
	Inicio  string `json:"inicio"`
	Fin     string `json:"fin"`
}

type PeriodoAcademico struct {
	IDPeriodoAcad string `gorm:"primaryKey" json:"id_periodo_acad"`
	YearAcad      int    `json:"year_acad"`
	Periodo       int    `json:"periodo"`
}

func (PeriodoAcademico) TableName() string {
	return "periodo_academico"
}
