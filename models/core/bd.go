package core

import "time"

type Usuario struct {
	IDUser string `gorm:"column:id_user;size:10;primaryKey"`
	Nombre string `gorm:"column:nombre"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

type Docente struct {
	IDDocente    int    `gorm:"column:id_docente;primaryKey"`
	ClaveDocente string `gorm:"column:clave_docente;size:80"`
	IDUser       string `gorm:"column:id_user;size:10;unique"`
}

func (Docente) TableName() string {
	return "docente"
}

type Estudiante struct {
	CodigoEstudiante int    `gorm:"column:codigo_estudiante;primaryKey"`
	ClaveEstudiante  string `gorm:"column:clave_estudiante;size:80"`
	IDUser           string `gorm:"column:id_user;size:10;unique"`
}

func (Estudiante) TableName() string {
	return "estudiante"
}

type ConsejoFacultad struct {
	IDConsejo    string `gorm:"column:id_consejo;primaryKey"`
	ClaveConsejo string `gorm:"column:clave_consejo;size:80"`
	IDUser       string `gorm:"column:id_user;size:10;unique"`
	IDFacultad   int    `gorm:"column:id_facultad"`
}

func (ConsejoFacultad) TableName() string {
	return "consejo_facultad"
}

type SecretarioAcademico struct {
	IDAcademico    string `gorm:"column:id_academico;primaryKey"`
	ClaveAcademico string `gorm:"column:clave_academico;size:80"`
	Correo         string `gorm:"column:correo"`
	IDUser         string `gorm:"column:id_user;size:10;unique"`
	IDFacultad     int    `gorm:"column:id_facultad"`
}

func (SecretarioAcademico) TableName() string {
	return "secretario_academico"
}

type SecretarioTecnico struct {
	IDSecret        string `gorm:"column:id_secret;primaryKey"`
	ClaveSecretario string `gorm:"column:clave_secretario;size:80"`
	Correo          string `gorm:"column:correo"`
	IDUser          string `gorm:"column:id_user;size:10;unique"`
}

func (SecretarioTecnico) TableName() string {
	return "secretario_tecnico"
}

type Cursos struct {
	IDCurso     int    `gorm:"column:id_curso;primaryKey"`
	NombreCurso string `gorm:"column:nombre_curso;size:45"`
	IDPrograma  int    `gorm:"column:id_programa"`
}

func (Cursos) TableName() string {
	return "curso"
}

type Programa struct {
	IdPrograma  int    `gorm:"column:id_programa;primaryKey"`
	NomPrograma string `gorm:"column:nombre_programa;size:45"`
	IdFacultad  int    `gorm:"column:id_facultad"`
}

func (Programa) TableName() string {
	return "programa"
}

type Facultad struct {
	IdFacultad     int    `gorm:"column:id_facultad;primaryKey"`
	NombreFacultad string `gorm:"column:nombre_facultad;size:55"`
}

func (Facultad) TableName() string {
	return "facultad"
}

type Ejerciendo struct {
	IDejerciendo int    `gorm:"column:id_ejerciendo;primaryKey;autoIncrement"`
	IDdocente    int    `gorm:"column:id_docente"`
	IDcurso      int    `gorm:"column:id_curso"`
	IDperiodo    string `gorm:"column:id_periodo_acad;size:10"`
	IDTipo       string `gorm:"column:id_tipo;size:10"`
}

func (Ejerciendo) TableName() string {
	return "ejerciendo"
}

type Cursando struct {
	IDejerciendo int    `gorm:"column:id_cursando;primaryKey;autoIncrement"`
	IDcurso      int    `gorm:"column:id_curso"`
	CodigoEst    int    `gorm:"column:codigo_estudiante"`
	IDperiodo    string `gorm:"column:id_periodo_acad;size:10"`
}

func (Cursando) TableName() string {
	return "cursando"
}

type Tipos struct {
	IDTipo      string `gorm:"column:id_tipo;size:10;primaryKey"`
	NombreTipo  string `gorm:"column:nombre_tipo;size:80"`
	Descripcion string `gorm:"column:descripcion"`
}

func (Tipos) TableName() string {
	return "tipo"
}

type Formatos struct {
	IdFormato     string `gorm:"column:id_formato;primaryKey"`
	NombreFormato string `gorm:"column:nombre_formato"`
	Descripcion   string `gorm:"column:descripcion"`
	Peso          int    `gorm:"column:peso"`
}

func (Formatos) TableName() string {
	return "formatos"
}

type Criterios struct {
	IdCriterio      string `gorm:"column:id_criterio;primaryKey"`
	Nombre_criterio string `gorm:"column:nombre_criterio"`
	IdFormato       string `gorm:"column:id_formato"`
}

func (Criterios) TableName() string {
	return "criterios_evaluacion"
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

func (Evaluacion) TableName() string {
	return "evaluacion"
}

type PeriodoEvaluacion struct {
	IDPeriodoEvl  string `gorm:"primaryKey" json:"id_periodo_evl"`
	FechaInicio   string `json:"fecha_inicio"`
	FechaFinal    string `json:"fecha_final"`
	IDPeriodoAcad string `json:"id_periodo_acad"`
}

func (PeriodoEvaluacion) TableName() string {
	return "periodo_evaluacion"
}

type PeriodoAcademico struct {
	IDPeriodoAcad string `gorm:"primaryKey" json:"id_periodo_acad"`
	YearAcad      int    `gorm:"column:year_acad" json:"year_acad"`
	Periodo       int    `gorm:"column:periodo" json:"periodo"`
	Fechainicial  string `gorm:"column:fecha_inicial" json:"fecha_inicial"`
	Fechafinal    string `gorm:"column:fecha_final" json:"fecha_final"`
}

func (PeriodoAcademico) TableName() string {
	return "periodo_academico"
}

type PDFHash struct {
	ID            uint      `gorm:"primaryKey"`
	NombreDocente string    `gorm:"type:varchar(255);not null"`
	Fecha         time.Time `gorm:"default:current_timestamp"`
	Hash          string    `gorm:"type:varchar(255);not null"`
}

func (PDFHash) TableName() string {
	return "pdf_hashes"
}
