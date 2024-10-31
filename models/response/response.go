package response

type Input struct {
	Switch string `json:"switch"`
	Estado bool   `json:"estado"`
}

type ReportePrograma struct {
	NombrePrograma string `json:"nombre_programa"`
}
