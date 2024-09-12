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
	r.GET("/periodoAcactivo", routers.PeriodoACactivo)
	r.PUT("/editarperiodo/:id", routers.Editarperiodo)
	r.GET("/cursos_facultad", routers.Curso_facultad)
	r.GET("/curso/:id", routers.Get_Curso)
	r.GET("/docentes", routers.Get_Docentes)
	r.POST("/asignar_docente", routers.AsignarEjerciendo)
	r.GET("/cursos_asignados", routers.Get_Cursos_asignados)
	r.POST("/docentes_asignados", routers.Get_Docentes_materia)
	//r.POST("/registro", middlewares.ChequeoBD(), routers.Registro)
}
