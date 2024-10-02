package bd

import (
	"log"
	"webservice/models"
)

func Get_Consejo(nom string) (models.ConsejoFacultad, error) {
	db := DBconexion()

	var miembro models.ConsejoFacultad

	err := db.Table("consejo_facultad").
		Joins("inner join usuarios on usuarios.id_user = consejo_facultad.id_user").
		Where("usuarios.nombre = ?", nom).
		Scan(&miembro).Error

	if err != nil {
		log.Fatalf("Error al obtener los docentes con usuario: %v", err)
		return models.ConsejoFacultad{}, err
	} else {
		return miembro, nil
	}
}

func Get_Docentes_facultad(consejo models.ConsejoFacultad) ([]string, error) {

	var docentesCursos []models.DocenteCurso

	periodoActivo, err := GetPeriodoAcActivo()
	if err != nil {
		return nil, err
	}

	ejerciendo, err := GetCursosByPeriodo(periodoActivo.IDPeriodoAcad)
	if err != nil {
		return nil, err
	}

	cursosFacultad, err := GetCursosPorFacultad(consejo.IDFacultad)
	if err != nil {
		return nil, err
	}

	var ejerciendoFiltrado []models.Ejerciendo
	for _, ejerce := range ejerciendo {
		for _, curso := range cursosFacultad {
			if ejerce.IDcurso == curso.IDCurso {
				ejerciendoFiltrado = append(ejerciendoFiltrado, ejerce)
				break
			}
		}
	}

	docentesCursos, err = GetDocentesYCursosByEjerciendo(ejerciendoFiltrado)

	uniqueDocentesMap := make(map[string]bool)

	for _, docenteCurso := range docentesCursos {
		uniqueDocentesMap[docenteCurso.NombreDocente] = true
	}

	uniqueDocentes := []string{}
	for docente := range uniqueDocentesMap {
		uniqueDocentes = append(uniqueDocentes, docente)
	}

	log.Println("Docentes :", uniqueDocentes)

	if err != nil {
		return nil, err
	}

	return uniqueDocentes, nil
}
