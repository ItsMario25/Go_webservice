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
	r.GET("/criterios_estudiante", routers.Get_Criterios_estudiante)
	r.POST("/evaluacion_estudiante", routers.Guardar_evaluacion)
	r.GET("/cursos_ejerciendo", routers.Get_Docentes_curso)
	r.GET("/criterios_docente", routers.Get_Criterios_docente)
	r.POST("/autoevaluacion_docente", routers.Guardar_autoevaluacion)
	r.GET("/docentes_facultad", routers.Get_Docentes_facultad)
	r.GET("/criterios_facultad", routers.Get_Criterios_facultad)
	r.POST("/evaluacion_facultad", routers.Guardar_evaluacion_facultad)
	r.GET("/docentes_evaluados", routers.Get_docentes_evaluados)
	r.GET("/cursos_evaluados", routers.Get_cursos_evaluados)
	r.GET("/ejerciendo", routers.ValidarEjerciendo)
	r.GET("/tipos", routers.Get_Tipos)
	r.GET("/reportes", routers.Historial_individual)
	r.POST("/reporte_individual", routers.Reporte_individual)
	r.POST("/reporte_general", routers.Reporte_general)
	r.POST("/validar_token_email", routers.Validar_Token)
	r.GET("/periodos_facultad", routers.Get_Periodos_facultad)
	r.GET("/switch_seguridad", routers.GetSwitches)
	r.POST("/switch_seguridad", routers.UpdateSwitch)
	//r.POST("/registro", middlewares.ChequeoBD(), routers.Registro)
}
