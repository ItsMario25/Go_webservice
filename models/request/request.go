package request

type GetAsignacion struct {
	Idcurso   string `json:"id_curso"`
	IDdocente string `json:"id_docente"`
	IDtipo    string `json:"id_tipo"`
}

type FormatoEvaluacion struct {
	NombreCurso     string         `json:"nombreCurso"`
	NombreDocente   string         `json:"nombreDocente"`
	NombreEvaluador string         `json:"nombreEvaluador"`
	Respuestas      map[int]string `json:"respuestas"` // Map donde la clave es el índice del criterio y el valor es la respuesta
}

type EvaluacionDocente struct {
	NombreDocente string         `json:"nombreDocente"`
	Respuestas    map[int]string `json:"respuestas"`
}

type FormatoEvaluacionFacultad struct {
	NombreDocente   string         `json:"nombreDocente"`
	NombreEvaluador string         `json:"nombreEvaluador"`
	Respuestas      map[int]string `json:"respuestas"` // Map donde la clave es el índice del criterio y el valor es la respuesta
}

type ReporteRequest struct {
	PeriodoAcademico string `json:"periodo_academico" binding:"required"`
	Vinculacion      string `json:"vinculacion" binding:"required"`
}

type ReportegRequest struct {
	PeriodoAcademico string `json:"periodo_academico" binding:"required"`
	Programanombre   string `json:"nombre_programa" binding:"required"`
}

type TokenRequest struct {
	ClientID string `json:"client_id"`
	Rol_us   string `json:"rol"`
}

type Credentials struct {
	Usuario    string `json:"usuario"`
	Contrasena string `json:"contrasena"`
	ClientID   string `json:"client_id"`
}

type Docentebdd struct {
	Idocente string `gorm:"column:id_docente" json:"id_docente"`
	Nombre   string `gorm:"column:nombre" json:"nombre"`
}
