package models

import "gopkg.in/mgo.v2/bson"

// Review models the data to review a Product
type Review struct {
	ID   bson.ObjectId `json:"id" bson:"id"`
	Text string        `json:"text" bson:"text"`
}
