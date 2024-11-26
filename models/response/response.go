package response

type Input struct {
	Switch string `json:"switch"`
	Estado bool   `json:"estado"`
}

type ReportePrograma struct {
	NombrePrograma string `json:"nombre_programa"`
}

type PDFHash struct {
	NombreDocente string `json:"nombre_docente"`
	Fecha         string `json:"fecha"`
}
