package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"webservice/middlewares"
	"webservice/routers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	// Define una ruta para la página de inicio
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bienvenido al servidor de mi-paquete",
		})
	})
	// Ruta para la página de login
	r.GET("/login", func(c *gin.Context) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		filepath := filepath.Join(dir, "templates", "login.html")
		fileContent, err := os.ReadFile(filepath)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error al leer archivo: "+err.Error())
			return
		}
		// Renderizar contenido del archivo en HTML
		c.Data(http.StatusOK, "text/html; charset=utf-8", fileContent)
	})

	r.POST("/registro", middlewares.ChequeoBD(), routers.Registro)

	return r
}
