package models

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
