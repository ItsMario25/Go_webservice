package routers

import (
	"log"
	"net/http"
	"webservice/bd"
	"webservice/models/core"
	"webservice/models/response"

	"github.com/gin-gonic/gin"
)

func GetSwitches(c *gin.Context) {

	configuraciones := bd.Confi_segur()
	// Enviar la configuración en formato JSON
	log.Println(configuraciones)
	c.JSON(http.StatusOK, gin.H{
		"multifactor":      getEstado("multifactor", configuraciones),
		"copia_controlada": getEstado("copia_controlada", configuraciones),
	})
}

func getEstado(switchName string, configs []core.ConfigSeguridad) bool {
	for _, config := range configs {
		if config.SwitchName == switchName {
			return config.Estado
		}
	}
	return false
}

func UpdateSwitch(c *gin.Context) {
	var input response.Input

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	bd.Update_config(input.Switch, input.Estado)

	c.JSON(http.StatusOK, gin.H{"status": "cambio guardado"})
}
