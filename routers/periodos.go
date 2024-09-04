package routers

import (
	"log"
	"net/http"
	"webservice/bd"

	"github.com/gin-gonic/gin"
)

func GetPeriodos(c *gin.Context) {
	consulta, err := bd.GetPeriod()

	if err != nil {
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, consulta)
	}
}
