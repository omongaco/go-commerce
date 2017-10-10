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
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Slug        string        `json:"slug"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	OrderStatus string        `json:"orderStatus"`
	Products    []*Products   `json:"products"`
	UserID      *User         `json:"userId"`
}

// Orders are a slice of Order
type Orders []Order
