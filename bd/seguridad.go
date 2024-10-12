package bd

import (
	"log"
	"webservice/models/core"
)

func SavePDFHash(nombreDocente, hash string) error {

	db := DBconexion()

	pdfHash := core.PDFHash{
		NombreDocente: nombreDocente,
		Hash:          hash,
	}

	if err := db.Create(&pdfHash).Error; err != nil {
		log.Println("Error al guardar el hash del PDF:", err)
		return err
	}

	return nil
}

func Confi_segur() []core.ConfigSeguridad {
	db := DBconexion()
	var configuraciones []core.ConfigSeguridad
	if err := db.Find(&configuraciones).Error; err != nil {
		return []core.ConfigSeguridad{}
	}

	return configuraciones
}

func Get_multifactor() bool {
	db := DBconexion()
	var configuraciones core.ConfigSeguridad
	// Corrección: agregar comillas simples para hacer la comparación con una cadena literal
	if err := db.Where("switch_name = ?", "multifactor").First(&configuraciones).Error; err != nil {
		return false
	}

	return configuraciones.Estado
}

func GetConfigSeguridad(switchName string) (core.ConfigSeguridad, error) {
	db := DBconexion()
	var config core.ConfigSeguridad
	if err := db.Where("switch_name = ?", switchName).First(&config).Error; err != nil {
		return config, err
	}
	return config, nil
}

func Update_config(name string, bb bool) {
	db := DBconexion()

	var config core.ConfigSeguridad
	if err := db.Where("switch_name = ?", name).First(&config).Error; err != nil {
		// Si no existe, lo creamos
		config.SwitchName = name
		config.Estado = bb
		db.Create(&config)
	} else {
		// Si existe, lo actualizamos
		config.Estado = bb
		log.Println(config)
		db.Save(&config)
	}
}
