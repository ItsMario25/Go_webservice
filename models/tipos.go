package models

type Tipos struct {
	IDTipo      string `gorm:"column:id_tipo;size:10;primaryKey"`
	NombreTipo  string `gorm:"column:nombre_tipo;size:80"`
	Descripcion string `gorm:"column:descripcion"`
}

func (Tipos) TableName() string {
	return "tipo"
}
