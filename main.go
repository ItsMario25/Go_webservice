package main

import (
	"log"
	"time"
	"webservice/bd"
	"webservice/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//utilidades.Crear_pdf()

	bd.InitDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handlers.InitRoutes(r)
	log.Println("Ejecutando servicio en http://localhost:8080/login")
	r.Run(":8080")

}
