package models

import "gopkg.in/mgo.v2/bson"

/*

Product is the struct which will model the data to MongoDB

ID is the ObjectId that is required by MongoDB, which is automatically added
Name is the name of that Product
Slug is a unique identifier, allowing a Product to have SEO-friendly URIs
Price is a floating point integer

*/
type Product struct {
	ID    bson.ObjectId `json:"_id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Slug  string        `json:"slug" bson:"slug"`
	Price float64       `json:"price" bson:"price"`
}

// Products are a slice of Product
type Products []Product
