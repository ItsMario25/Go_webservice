package handlers

import (
	"webservice/middlewares"
	"webservice/routers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", routers.GetLogin)
	r.GET("/login", routers.GetLogin)
	r.GET("/registro", routers.GetRegistro)
	r.POST("/registro", middlewares.ChequeoBD(), routers.Registro)

	return r
}
