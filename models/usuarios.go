package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Usuario struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre   string             `bson:"nombre" json:"nombre,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
}

type Logg struct {
	Email    string
	Password string
}
