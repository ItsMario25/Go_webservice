package models

type Estudiante struct {
	CodigoEstudiante int    `gorm:"column:codigo_estudiante"`
	ClaveEstudiante  string `gorm:"column:clave_estudiante"`
}
