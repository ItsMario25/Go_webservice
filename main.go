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
		AllowOrigins:     []string{"https://localhost:3000"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "Referer"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handlers.InitRoutes(r)
	log.Println("Ejecutando servicio en https://localhost:8080/")
	//r.Run(":8080")
	err := r.RunTLS(":8080", "cert.pem", "key.pem")
	if err != nil {
		log.Fatal("No se pudo iniciar el servidor HTTPS:", err)
	}

}
