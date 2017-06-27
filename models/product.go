package models

import "gopkg.in/mgo.v2/bson"

// Product models our data for the database
type Product struct {
	ID            bson.ObjectId `json:"_id" bson:"_id"`
	Name          string        `json:"name" bson:"name"`
	Slug          string        `json:"slug" bson:"slug"`
	Description   string        `json:"description" bson:"description"`
	FeaturedImage string        `json:"featured_image" bson:"featured_image"`
	Gallery       []string      `json:"gallery" bson:"gallery"`
	Price         float64       `json:"price" bson:"price"`
	InStock       bool          `json:"in_stock" bson:"in_stock"`
	Category      *Category     `json:"category" bson:"category"`
}

// Products are a slice of Product
type Products []Product
