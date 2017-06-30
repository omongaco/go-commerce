package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Product models our data for the database
type Product struct {
	ID            bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt     time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at" bson:"updated_at"`
	Name          string        `json:"name" bson:"name"`
	Slug          string        `json:"slug" bson:"slug"`
	Description   string        `json:"description" bson:"description"`
	FeaturedImage string        `json:"featured_image" bson:"featured_image"`
	Gallery       []string      `json:"gallery" bson:"gallery"`
	QuantityLeft  float64       `json:"quantity_left" bson:"quantity_left"`
	InStock       bool          `json:"in_stock" bson:"in_stock"`
	Price         float64       `json:"price" bson:"price"`
	Subtotal      float64       `json:"subtotal" bson:"subtotal"`
	Total         float64       `json:"total" bson:"total"`
	Reviews       []*Review     `json:"reviews" bson:"reviews"`
}

// Products are a slice of Product
type Products []Product
