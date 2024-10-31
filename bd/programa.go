package bd

import (
	"webservice/models/core"
	"webservice/models/response"
)

func GetProgramasPorFacultad(idFacultad int) ([]response.ReportePrograma, error) {
	db := DBconexion()
	var programas []response.ReportePrograma

	err := db.Table("programa").
		Select("programa.nombre_programa").
		Where("programa.id_facultad = ?", idFacultad).
		Scan(&programas).Error

	if err != nil {
		return nil, err
	}
	return programas, nil
}

func GerProgramaPorNombre(nombre string) (core.Programa, error) {
	db := DBconexion()

	var programas core.Programa
	err := db.Table("programa").
		Select("programa.id_programa").
		Where("programa.nombre_programa = ?", nombre).
		Scan(&programas).Error

	if err != nil {
		return core.Programa{}, err
	}
	return programas, nil

}
