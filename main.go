package main

import (
	"log"
	"webservice/bd"
	"webservice/handlers"

	"github.com/gin-contrib/cors"
)

func main() {
	if bd.CheckConnect() == 0 {
		log.Fatal("Sin conexion con la BD")
		return
	}
	r := handlers.InitRoutes()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	log.Println("Ejecutando servicio en http://localhost:8080/login")
	r.Run(":8080")
}
