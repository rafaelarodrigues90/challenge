package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User database
type User struct {
	ID			primitive.ObjectID				`bson:"_id"`
	FirstName	string				`bson:"first_name"`
	LastName	string				`bson:"last_name"`
	CPF 		int 				`bson:"cpf"` // como validar?
	Saldo		float32				`bson:"saldo"` // NÃO PODE SER NEGATIVO
	Email		string 				`bson:"email"` // como validar?
}

type Users []User