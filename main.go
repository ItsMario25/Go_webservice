package main

import (
	"log"
	"webservice/bd"
	"webservice/routes"

	"github.com/gin-contrib/cors"
)

func main() {
	if bd.CheckConnect() == 0 {
		log.Fatal("Sin conexion con la BD")
		return
	}
	r := routes.InitRoutes()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.Run(":8080")
}
