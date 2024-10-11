package core

type DocenteCurso struct {
	NombreDocente string `json:"nombre_docente"`
	NombreCurso   string `json:"nombre_curso"`
}

type PeriodoAc struct {
	Periodo string `json:"periodo"`
	Inicio  string `json:"inicio"`
	Fin     string `json:"fin"`
}

type EvaluacionReporte struct {
	PeriodoAcad string `json:"periodo_academico"`
	FechaFinal  string `json:"fecha_final"`
	TipoDocente string `json:"tipo_docente"`
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

type DocenteConUsuario struct {
	IdDocente int
	Nombre    string
}

type ConfigSeguridad struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	SwitchName string `gorm:"column:switch_name" json:"switch_name"`
	Estado     bool   `gorm:"column:estado" json:"estado"`
}

func (ConfigSeguridad) TableName() string {
	return "configuracion_seguridad"
}
