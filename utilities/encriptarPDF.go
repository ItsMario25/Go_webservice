package utilities

import (
	"crypto/sha256"
	"encoding/hex"
)

// Función para generar el hash del PDF
func GeneratePDFHash(pdfBytes []byte) string {
	hash := sha256.Sum256(pdfBytes)
	return hex.EncodeToString(hash[:])
}
