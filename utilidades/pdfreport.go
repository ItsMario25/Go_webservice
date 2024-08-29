package utilidades

import "github.com/jung-kurt/gofpdf/v2"

func Crear_pdf() {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 12)

	// Crear una tabla de 3 columnas y 5 filas
	header := []string{"Columna 1", "Columna 2", "Columna 3"}
	data := [][]string{
		{"Fila 1, Col 1", "Fila 1, Col 2", "Fila 1, Col 3"},
		{"Fila 2, Col 1", "Fila 2, Col 2", "Fila 2, Col 3"},
		{"Fila 3, Col 1", "Fila 3, Col 2", "Fila 3, Col 3"},
		{"Fila 4, Col 1", "Fila 4, Col 2", "Fila 4, Col 3"},
		{"Fila 5, Col 1", "Fila 5, Col 2", "Fila 5, Col 3"},
	}

	colWidths := []float64{60, 60, 60}

	for i, str := range header {
		pdf.CellFormat(colWidths[i], 7, str, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	for _, row := range data {
		for i, str := range row {
			pdf.CellFormat(colWidths[i], 7, str, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}

	// Guardar el archivo PDF
	fileName := "unprotected_table.pdf"
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		panic(err)
	}
}
