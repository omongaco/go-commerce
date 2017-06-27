package models

import "gopkg.in/mgo.v2/bson"

// Order models our data for the database
type Order struct {
	ID bson.ObjectId `json:"_id" bson:"_id"`
}

// Orders are a slice of Order
type Orders []Order
