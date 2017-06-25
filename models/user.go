package models

import "gopkg.in/mgo.v2/bson"

/*

User is the struct which will model the data to MongoDB

*/
type User struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	FirstName string        `json:"first_name" bson:"first_name"`
	LastName  string        `json:"last_name" bson:"last_name"`
	Email     string        `json:"email" bson:"last_name"`
	Password  string        `json:"password" bson:"password"`
	Phone     string        `json:"phone" bson:"phone"`
	Gender    string        `json:"gender" bson:"gender"`
	Address   string        `json:"address" bson:"address"`
	City      string        `json:"city" bson:"city"`
	State     string        `json:"state" bson:"state"`
	Zip       string        `json:"zip" bson:"zip"`
	Country   string        `json:"country" bson:"country"`
}

// Users are a slice of User
type Users []User