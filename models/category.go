package models

import "gopkg.in/mgo.v2/bson"

// Category models our data for the database
type Category struct {
	ID   bson.ObjectId `json:"_id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}
