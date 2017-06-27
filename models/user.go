package models

import "gopkg.in/mgo.v2/bson"

// User models our data for the database
type User struct {
	ID               bson.ObjectId `json:"_id" bson:"_id"`
	FirstName        string        `json:"first_name" bson:"first_name"`
	LastName         string        `json:"last_name" bson:"last_name"`
	Email            string        `json:"email" bson:"email"`
	Password         string        `json:"password" bson:"password"`
	Phone            string        `json:"phone" bson:"phone"`
	Gender           string        `json:"gender" bson:"gender"`
	Address          string        `json:"address" bson:"address"`
	City             string        `json:"city" bson:"city"`
	State            string        `json:"state" bson:"state"`
	Zip              string        `json:"zip" bson:"zip"`
	Country          string        `json:"country" bson:"country"`
	FavoriteProducts *Products     `json:"favorite_products" bson:"favorite_products"`
	BoughtProducts   *Products     `json:"bought_products" bson:"bought_products"`
	CartItems        *Products     `json:"cart_items" bson:"cart_items"`
}

// Users are a slice of User
type Users []User
