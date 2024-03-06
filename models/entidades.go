package models

type Curso struct {
	id_curso     int
	nombre_curso string
	id_programa  int
}

type Docente struct {
	cedula         int
	nombre_docente string
	id_rol         int
	id_tipo        int
	clave_docente  string
}

type Estudiante struct {
	codigo_estudiante int
	nombre_estudiante string
	id_rol            int
	Clave_estudiante  string
}

type Evaluacion struct {
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

type Facultad struct {
	id_facultad     int
	nombre_facultad string
}

type Programa struct {
	id_programa     int
	nombre_programa string
	id_facultad     int
}

type Roles struct {
	id_rol      int
	nombre_tipo string
}
