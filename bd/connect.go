package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = connectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://marioGo:Redstorm31@cluster1.blbpbqu.mongodb.net/?retryWrites=true&w=majority")

func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Conexion Exitosa con la DB Mongo")
	return client
}

func CheckConnect() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
