package models

type EvaluacionReporte struct {
	PeriodoAcad string `json:"periodo_academico"`
	FechaFinal  string `json:"fecha_final"`
	TipoDocente string `json:"tipo_docente"`
}

type ReporteRequest struct {
	PeriodoAcademico string `json:"periodo_academico" binding:"required"`
	Vinculacion      string `json:"vinculacion" binding:"required"`
}

type EvaluacionReporteEstudiante struct {
	IdEvaluacion      int    `gorm:"column:id_evaluacion;primaryKey;autoIncrement"`
	FechaDiligenciada string `gorm:"column:fecha_diligenciada"`
	Calificacion      string `gorm:"column:calificacion"`
	IdCriterio        string `gorm:"column:id_criterio"`
	IdUser            string `gorm:"column:id_user"`
	Curso             string `gorm:"column:id_curso"`
	NomCurso          string `gorm:"column:nombre_curso;size:45"`
}

type EvaluacionReporteDC struct {
	IdEvaluacion      int    `gorm:"column:id_evaluacion;primaryKey;autoIncrement"`
	FechaDiligenciada string `gorm:"column:fecha_diligenciada"`
	Calificacion      string `gorm:"column:calificacion"`
	Iduser            string `gorm:"column:id_user"`
}
