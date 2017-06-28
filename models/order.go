package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*
Order models our data for the database

ID is the ObjectID for a specific Order
Slug will
UserID will be that specific User's ID who is buying that Product
OrderStatus will be: paid, shipping, received or refunded
UserID is the ObjectId of the User that a specific Order belongs to

*/
type Order struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Slug        string        `json:"slug" bson:"slug"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
	UserID      *User         `json:"user_id" bson:"user_id"`
	OrderStatus string        `json:"order_status" bson:"order_status"`
}

// Orders are a slice of Order
type Orders []Order
