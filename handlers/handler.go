package handlers

import (
	"webservice/routers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/verificar", routers.ValidarLogin)
	r.POST("/validar_token", routers.ValidarToken)
	r.GET("/periodos_evl", routers.GetPeriodos)
	//r.GET("/login", routers.GetLogin)
	//r.GET("/registro", routers.GetRegistro)
	//r.POST("/registro", middlewares.ChequeoBD(), routers.Registro)
}
