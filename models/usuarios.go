package models

type Usuario struct {
	IDUser string `gorm:"column:id_user;primaryKey"`
	Nombre string `gorm:"column:nombre"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

type Docente struct {
	IDDocente    int    `gorm:"column:id_docente;primaryKey"`
	ClaveDocente string `gorm:"column:clave_docente"`
	IDUser       string `gorm:"column:id_user;unique"`
}

func (Docente) TableName() string {
	return "docente"
}

type Estudiante struct {
	CodigoEstudiante int    `gorm:"column:codigo_estudiante;primaryKey"`
	ClaveEstudiante  string `gorm:"column:clave_estudiante"`
	IDUser           string `gorm:"column:id_user;unique"`
}

func (Estudiante) TableName() string {
	return "estudiante"
}

type ConsejoFacultad struct {
	IDConsejo    string `gorm:"column:id_consejo;primaryKey"`
	ClaveConsejo string `gorm:"column:clave_consejo"`
	IDUser       string `gorm:"column:id_user;unique"`
	IDFacultad   int    `gorm:"column:id_facultad"`
}

func (ConsejoFacultad) TableName() string {
	return "consejo_facultad"
}

type SecretarioAcademico struct {
	IDAcademico    string `gorm:"column:id_academico;primaryKey"`
	ClaveAcademico string `gorm:"column:clave_academico"`
	IDUser         string `gorm:"column:id_user;unique"`
}

func (SecretarioAcademico) TableName() string {
	return "secretario_academico"
}

type SecretarioTecnico struct {
	IDSecret        string `gorm:"column:id_secret;primaryKey"`
	ClaveSecretario string `gorm:"column:clave_secretario"`
	IDUser          string `gorm:"column:id_user;unique"`
}

func (SecretarioTecnico) TableName() string {
	return "secretario_tecnico"
}

type DocenteConUsuario struct {
	IdDocente int
	Nombre    string
}
