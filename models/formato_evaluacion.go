package models

// Estructura para recibir la evaluación
type EvaluacionEstudiante struct {
	NombreCurso     string         `json:"nombreCurso"`
	NombreDocente   string         `json:"nombreDocente"`
	NombreEvaluador string         `json:"nombreEvaluador"`
	Respuestas      map[int]string `json:"respuestas"` // Map donde la clave es el índice del criterio y el valor es la respuesta
}
