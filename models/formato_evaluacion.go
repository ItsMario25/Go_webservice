package models

// Estructura para recibir la evaluación
type FormatoEvaluacion struct {
	NombreCurso     string         `json:"nombreCurso"`
	NombreDocente   string         `json:"nombreDocente"`
	NombreEvaluador string         `json:"nombreEvaluador"`
	Respuestas      map[int]string `json:"respuestas"` // Map donde la clave es el índice del criterio y el valor es la respuesta
}

type EvaluacionDocente struct {
	NombreCurso   string         `json:"nombreCurso"`
	NombreDocente string         `json:"nombreDocente"`
	Respuestas    map[int]string `json:"respuestas"`
}

type Evaluacion struct {
	IdEvaluacion      int    `gorm:"column:id_evaluacion;primaryKey;autoIncrement"`
	FechaDiligenciada string `gorm:"column:fecha_diligenciada"`
	Calificacion      string `gorm:"column:calificacion"`
	IdCriterio        string `gorm:"column:id_criterio"`
	IdUser            string `gorm:"column:id_user"`
	IdPeriodoEvl      string `gorm:"column:id_periodo_evl"`
	IdEjerciendo      int    `gorm:"column:id_ejerciendo"`
}

// Método para personalizar el nombre de la tabla
func (Evaluacion) TableName() string {
	return "evaluacion"
}
