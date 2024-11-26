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
	log.Println("Ejecutando servicio en http://localhost:8080/")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("No se pudo iniciar el servidor:", err)
	}

}


