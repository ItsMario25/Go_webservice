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
}

type Ejerciendo struct {
	IDejerciendo int    `gorm:"column:id_ejerciendo;primaryKey;autoIncrement"`
	IDdocente    int    `gorm:"column:id_docente"`
	IDcurso      int    `gorm:"column:id_curso"`
	IDperiodo    string `gorm:"column:id_periodo_acad;size:10"`
}

func (Ejerciendo) TableName() string {
	return "ejerciendo"
}
