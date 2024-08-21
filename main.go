package main

import (
	"log"
	"time"
	"webservice/handlers"

	"github.com/gin-contrib/cors"
)

func main() {
	/*
		if bd.CheckConnect() == 0 {
			log.Fatal("Sin conexion con la BD")
			return
		}
	*/
	r := handlers.InitRoutes()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	log.Println("Ejecutando servicio en http://localhost:8080/login")
	r.Run(":8080")

}
