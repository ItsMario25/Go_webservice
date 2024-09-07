package models

type Cursos struct {
	IDCurso     int    `gorm:"column:id_curso;primaryKey;autoIncrement"`
	NombreCurso string `gorm:"column:nombre_curso;size:45"`
	IDPrograma  int    `gorm:"column:id_programa"`
}

func (Cursos) TableName() string {
	return "curso"
}
