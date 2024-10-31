package bd

import (
	"fmt"
	"webservice/models/core"
	"webservice/models/request"
)

func Get_historial_ejerciendo(docente int) ([]core.Ejerciendo, error) {
	db := DBconexion()

	var ejerciendoID []core.Ejerciendo

	err := db.Table("ejerciendo").
		Select("id_ejerciendo, id_tipo, id_periodo_acad").
		Where("id_docente = ?", docente).
		Scan(&ejerciendoID).Error

	if err != nil {
		return []core.Ejerciendo{}, err
	} else {
		return ejerciendoID, nil
	}
}

func Get_reportes_semestre(docente int) ([]core.EvaluacionReporte, error) {
	db := DBconexion()
	var reportes []core.EvaluacionReporte

	err := db.Table("evaluacion").
		Select(
			"periodo_evaluacion.id_periodo_acad AS periodo_academico",
			"periodo_evaluacion.fecha_final AS fecha_final",
			"tipo.nombre_tipo AS tipo_docente").
		Joins("JOIN ejerciendo ON evaluacion.id_ejerciendo = ejerciendo.id_ejerciendo").
		Joins("JOIN periodo_evaluacion ON periodo_evaluacion.id_periodo_acad = ejerciendo.id_periodo_acad").
		Joins("JOIN tipo ON ejerciendo.id_tipo = tipo.id_tipo").
		Where("ejerciendo.id_docente = ?", docente).
		Where("periodo_evaluacion.fecha_final <= NOW()").
		Group("periodo_evaluacion.id_periodo_acad, periodo_evaluacion.fecha_final, tipo.nombre_tipo").
		Scan(&reportes).Error

	if err != nil {
		return []core.EvaluacionReporte{}, err
	}

	for i, reporte := range reportes {
		var idPeriodoAcad string

		// Segunda consulta para obtener el id_periodo_acad
		err := db.Table("periodo_evaluacion").
			Select("id_periodo_acad").
			Where("fecha_final = ?", reporte.FechaFinal).
			Scan(&idPeriodoAcad).Error

		if err != nil {
			return []core.EvaluacionReporte{}, err
		}
		reportes[i].PeriodoAcad = idPeriodoAcad
	}

	return reportes, nil
}

func Get_resultados_formato_consejo(periodo string, usuario string) ([]core.EvaluacionReporteDC, error) {
	db := DBconexion()
	var resultados []core.EvaluacionReporteDC

	err := db.Table("evaluacion").
		Select("evaluacion.id_evaluacion, evaluacion.fecha_diligenciada, evaluacion.calificacion, evaluacion.id_user").
		Joins("JOIN criterios_evaluacion ON evaluacion.id_criterio = criterios_evaluacion.id_criterio").
		Joins("JOIN formatos ON criterios_evaluacion.id_formato = formatos.id_formato").
		Joins("JOIN ejerciendo ON evaluacion.id_ejerciendo = ejerciendo.id_ejerciendo").
		Joins("JOIN docente ON ejerciendo.id_docente = docente.id_docente").
		Where("formatos.id_formato = ?", "F01").
		Where("evaluacion.id_periodo_evl = ?", periodo).
		Where("docente.id_user = ?", usuario).
		Scan(&resultados).Error

	if err != nil {
		return nil, err
	}

	return resultados, nil
}

func Get_resultados_formato_estudiante(periodo string, usuario string) ([]core.EvaluacionReporteEstudiante, error) {
	db := DBconexion()
	var resultados []core.EvaluacionReporteEstudiante

	err := db.Table("evaluacion").
		Select("evaluacion.id_evaluacion, evaluacion.fecha_diligenciada, evaluacion.calificacion, evaluacion.id_user, curso.id_curso, curso.nombre_curso").
		Joins("JOIN criterios_evaluacion ON evaluacion.id_criterio = criterios_evaluacion.id_criterio").
		Joins("JOIN formatos ON criterios_evaluacion.id_formato = formatos.id_formato").
		Joins("JOIN ejerciendo ON evaluacion.id_ejerciendo = ejerciendo.id_ejerciendo").
		Joins("JOIN docente ON ejerciendo.id_docente = docente.id_docente").
		Joins("JOIN curso ON ejerciendo.id_curso = curso.id_curso").
		Where("formatos.id_formato = ?", "F02").
		Where("evaluacion.id_periodo_evl = ?", periodo).
		Where("docente.id_user = ?", usuario).
		Scan(&resultados).Error

	if err != nil {
		return nil, err
	}
	return resultados, nil
}

func Get_resultados_formato_docente(periodo string, usuario string) ([]core.EvaluacionReporteDC, error) {
	db := DBconexion()
	var resultados []core.EvaluacionReporteDC

	err := db.Table("evaluacion").
		Select("evaluacion.id_evaluacion, evaluacion.fecha_diligenciada, evaluacion.calificacion, evaluacion.id_user").
		Joins("JOIN criterios_evaluacion ON evaluacion.id_criterio = criterios_evaluacion.id_criterio").
		Joins("JOIN formatos ON criterios_evaluacion.id_formato = formatos.id_formato").
		Where("evaluacion.id_periodo_evl = ? AND formatos.id_formato = ? AND evaluacion.id_user = ? ", periodo, "F03", usuario).
		Scan(&resultados).Error

	if err != nil {
		return nil, err
	}
	return resultados, nil
}

func CalcularCalificacionTotal(resultados []core.EvaluacionReporteDC) float64 {
	total := 0
	for _, resultado := range resultados {
		switch resultado.Calificacion {
		case "nunca":
			total += 1
		case "algunas_veces":
			total += 3
		case "casi_siempre":
			total += 4
		case "siempre":
			total += 5
		}
	}
	if len(resultados) > 0 {
		return float64(total) / float64(len(resultados))
	}
	return 0
}

func CalcularCalificacionPorCurso(resultados []core.EvaluacionReporteEstudiante) (map[string]string, map[string]int) {
	calificacionesPorCurso := make(map[string]int)
	conteoPorCurso := make(map[string]int)
	estudiantePorCurso := make(map[string]int)

	for _, resultado := range resultados {
		switch resultado.Calificacion {
		case "nunca":
			calificacionesPorCurso[resultado.Curso] += 1
		case "algunas_veces":
			calificacionesPorCurso[resultado.Curso] += 3
		case "casi_siempre":
			calificacionesPorCurso[resultado.Curso] += 4
		case "siempre":
			calificacionesPorCurso[resultado.Curso] += 5
		}
		conteoPorCurso[resultado.Curso]++
		estudiantePorCurso[resultado.NomCurso]++
	}

	// Calcular el promedio por curso
	promediosPorCurso := make(map[string]string)
	for curso, total := range calificacionesPorCurso {
		valor := float64(total) / float64(conteoPorCurso[curso])
		promediosPorCurso[curso] = fmt.Sprintf("%.2f", valor)
	}

	for curso, total := range estudiantePorCurso {
		estudiantePorCurso[curso] = total / 23
	}

	return promediosPorCurso, estudiantePorCurso
}

func GetDocentesPorProgramaYPeriodo(idPrograma int, idPeriodoEvl string) ([]request.Docentebdd, error) {
	db := DBconexion()
	var docentes []request.Docentebdd

	err := db.Table("docente").
		Select("DISTINCT docente.id_docente, usuarios.nombre").
		Joins("JOIN usuarios ON docente.id_user = usuarios.id_user").
		Joins("JOIN ejerciendo ON docente.id_docente = ejerciendo.id_docente").
		Joins("JOIN curso ON ejerciendo.id_curso = curso.id_curso").
		Joins("JOIN programa ON curso.id_programa = programa.id_programa").
		Joins("JOIN periodo_evaluacion ON ejerciendo.id_periodo_acad = periodo_evaluacion.id_periodo_acad").
		Where("programa.id_programa = ? AND periodo_evaluacion.id_periodo_evl = ?", idPrograma, idPeriodoEvl).
		Scan(&docentes).Error

	if err != nil {
		return nil, err
	}
	return docentes, nil
}

func Get_resultados_consejo_general(idPeriodo string, idFacultad int) ([]core.EvaluacionReporteDC, error) {
	db := DBconexion()
	var resultados []core.EvaluacionReporteDC

	err := db.Table("evaluacion").
		Select("evaluacion.id_evaluacion, evaluacion.fecha_diligenciada, evaluacion.calificacion, docente.id_user, docente.nombre").
		Joins("JOIN criterios_evaluacion ON evaluacion.id_criterio = criterios_evaluacion.id_criterio").
		Joins("JOIN formatos ON criterios_evaluacion.id_formato = formatos.id_formato").
		Joins("JOIN ejerciendo ON evaluacion.id_ejerciendo = ejerciendo.id_ejerciendo").
		Joins("JOIN docente ON ejerciendo.id_docente = docente.id_docente").
		Where("formatos.id_formato = ? AND evaluacion.id_periodo_evl = ? AND docente.id_facultad = ?", "F01", idPeriodo, idFacultad).
		Scan(&resultados).Error

	if err != nil {
		return nil, err
	}
	return resultados, nil
}

func Get_resultados_estudiante_general(idPeriodo string, idFacultad string) ([]core.EvaluacionReporteEstudiante, error) {
	db := DBconexion()
	var resultados []core.EvaluacionReporteEstudiante

	err := db.Table("evaluacion").
		Select("evaluacion.id_evaluacion, evaluacion.fecha_diligenciada, evaluacion.calificacion, docente.id_user, docente.nombre, curso.id_curso, curso.nombre_curso").
		Joins("JOIN criterios_evaluacion ON evaluacion.id_criterio = criterios_evaluacion.id_criterio").
		Joins("JOIN formatos ON criterios_evaluacion.id_formato = formatos.id_formato").
		Joins("JOIN ejerciendo ON evaluacion.id_ejerciendo = ejerciendo.id_ejerciendo").
		Joins("JOIN docente ON ejerciendo.id_docente = docente.id_docente").
		Joins("JOIN curso ON ejerciendo.id_curso = curso.id_curso").
		Joins("JOIN programa ON curso.id_programa = programa.id_programa").
		Where("formatos.id_formato = ? AND evaluacion.id_periodo_evl = ? AND docente.id_facultad = ?", "F02", idPeriodo, idFacultad).
		Scan(&resultados).Error

	if err != nil {
		return nil, err
	}
	return resultados, nil
}

func Get_resultados_autoevaluacion_general(idPeriodo string, idFacultad string) ([]core.EvaluacionReporteDC, error) {
	db := DBconexion()
	var resultados []core.EvaluacionReporteDC

	err := db.Table("evaluacion").
		Select("evaluacion.id_evaluacion, evaluacion.fecha_diligenciada, evaluacion.calificacion, docente.id_user, docente.nombre").
		Joins("JOIN criterios_evaluacion ON evaluacion.id_criterio = criterios_evaluacion.id_criterio").
		Joins("JOIN formatos ON criterios_evaluacion.id_formato = formatos.id_formato").
		Joins("JOIN docente ON evaluacion.id_user = docente.id_user").
		Where("formatos.id_formato = ? AND evaluacion.id_periodo_evl = ? AND docente.id_facultad = ?", "F03", idPeriodo, idFacultad).
		Scan(&resultados).Error

	if err != nil {
		return nil, err
	}
	return resultados, nil
}
