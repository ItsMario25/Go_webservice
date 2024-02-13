package models

type curso struct {
	id_curso     int
	nombre_curso string
	id_programa  int
}

type docente struct {
	cedula         int
	nombre_docente string
	id_rol         int
	id_tipo        int
	clave_docente  string
}

type estudiante struct {
	codigo_estudiante int
	nombre_estudiante string
	id_rol            int
	clave_estudiante  string
}

type evaluacion struct {
	id_evaluacion      int
	periodo_evaluacion int
	fecha_diligenciada string
	year_evaluacion    int
	calificacion       string
	cedula             int
	id_curso           int
	codigo_estudiante  int
	id_criterio        string
}

type facultad struct {
	id_facultad     int
	nombre_facultad string
}

type programa struct {
	id_programa     int
	nombre_programa string
	id_facultad     int
}

type roles struct {
	id_rol      int
	nombre_tipo string
}
