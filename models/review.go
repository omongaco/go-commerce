package models

import "gopkg.in/mgo.v2/bson"

// Review models the data to review a Product
type Review struct {
	ID   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Text string        `json:"text" bson:"text"`
}

// Reviews are a slice of Review
type Reviews []Review
