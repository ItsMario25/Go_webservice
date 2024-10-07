package main

import (
	"log"
	"webservice/bd"
	"webservice/handlers"
	"webservice/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	bd.InitDB()

	r := gin.Default()

	// Aplicar el middleware de CORS
	r.Use(middlewares.CORSConfig())

	handlers.InitRoutes(r)
	log.Println("Ejecutando servicio en https://localhost:8080/")

	err := r.RunTLS(":8080", "cert.pem", "key.pem")
	if err != nil {
		log.Fatal("No se pudo iniciar el servidor HTTPS:", err)
	}

}
