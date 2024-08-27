package main

import (
	"log"
	"time"
	"webservice/bd"
	"webservice/handlers"

	"github.com/gin-contrib/cors"
)

func main() {
	bd.InitDB()
	r := handlers.InitRoutes()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	log.Println("Ejecutando servicio en http://localhost:8080/login")
	r.Run(":8080")

}
