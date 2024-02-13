package bd

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = connectDBnoSQ()
var clientOptions = options.Client().ApplyURI("mongodb+srv://marioGo:Redstorm31@cluster1.blbpbqu.mongodb.net/?retryWrites=true&w=majority")

func connectDBnoSQ() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Conexion Exitosa con la DB Mongo")
	return client
}
func CheckConnectnoSQ() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

var sqlDB = connectDB()

func connectDB() *sql.DB {
	dsn := "postgres://postgres:super1234@127.0.0.1:5432/securetrustdb?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexion con la base de datos realizada")
	return db
}

func CheckConnect() int {
	err := sqlDB.Ping()
	if err != nil {
		return 0
	}
	return 1
}
