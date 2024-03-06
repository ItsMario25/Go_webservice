package bd

import (
	"context"
	"time"
	"webservice/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegistro(u models.Estudiante) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twittor")
	col := db.Collection("proyect_go")

	u.Clave_estudiante, _ = EncriptarPassword(u.Clave_estudiante)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
