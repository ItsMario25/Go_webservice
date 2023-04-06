package routers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"webservice/models"

	"github.com/gin-gonic/gin"
)

/*Registro es la funcion para crear en la BD el registro de usuario */
func Registro(c *gin.Context) {
	var t models.Usuario

	if err := c.Bind(&t); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error en los datos recibidos " + err.Error(),
		})
		return
	}

	if len(t.Email) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "El email de usuario es requerido",
		})
		return
	}

	if len(t.Password) < 6 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Debe especificar una contraseña de al menos 6 caracteres",
		})
		return
	}

	/*c.JSON(http.StatusCreated, gin.H{
		"message":  "Usuario registrado exitosamente",
		"Usuario":  t.Nombre,
		"Password": t.Password,
		"Email":    t.Email,
	})*/

	GetIndex(c)
}

func GetIndex(c *gin.Context) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filepath := filepath.Join(dir, "templates", "index.html")
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error al leer archivo: "+err.Error())
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", fileContent)
}

func GetLogin(c *gin.Context) {
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
	c.Data(http.StatusOK, "text/html; charset=utf-8", fileContent)
}

func GetRegistro(c *gin.Context) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filepath := filepath.Join(dir, "templates", "registro.html")
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error al leer archivo: "+err.Error())
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", fileContent)
}
