package models

type Cursos struct {
	IDCurso     int    `gorm:"column:id_curso;primaryKey"`
	NombreCurso string `gorm:"column:nombre_curso;size:45"`
	IDPrograma  int    `gorm:"column:id_programa"`
}

func (Cursos) TableName() string {
	return "curso"
}

type GetAsignacion struct {
	Idcurso   string `json:"id_curso"`
	IDdocente string `json:"id_docente"`
	IDtipo    string `json:"id_tipo"`
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

type DocenteCurso struct {
	NombreDocente string `json:"nombre_docente"`
	NombreCurso   string `json:"nombre_curso"`
}
