package handlers

import (
	"webservice/routers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/verificar", routers.ValidarLogin)
	r.POST("/validar_token", routers.ValidarToken)
	r.GET("/periodos_evl", routers.GetPeriodos)
	r.POST("/cargarperiodo", routers.CargarPeriodo)
	r.GET("/periodoactivo", routers.PeriodoActivo)
	r.PUT("/editarperiodo/:id", routers.Editarperiodo)
	r.GET("/cursos_facultad", routers.Curso_facultad)
	r.GET("/curso/:id", routers.Get_Curso)
	r.GET("/docentes", routers.Get_Docentes)
	//r.POST("/registro", middlewares.ChequeoBD(), routers.Registro)
}
