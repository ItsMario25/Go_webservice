package bd

import (
	"context"
	"fmt"
	"time"
	"webservice/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoExistencia(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twittor")
	col := db.Collection("proyect_go")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := fmt.Sprintf("%x", resultado.ID)
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
