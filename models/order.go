package models

import "gopkg.in/mgo.v2/bson"

/*
Order models our data for the database

ID is the ObjectID assigned to a User at sign-up
UserID will be that specific User's ID who is buying that Product
OrderStatus will be: unpaid, paid, shipping, or received

*/
type Order struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	UserID      *User         `json:"user_id" bson:"user_id"`
	OrderStatus string        `json:"order_status" bson:""`
}

// Orders are a slice of Order
type Orders []Order
