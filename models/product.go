package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Product models our data for the database
type Product struct {
	ID            bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
	Name          string        `json:"name"`
	Slug          string        `json:"slug"`
	Description   string        `json:"description"`
	FeaturedImage string        `json:"featuredImage"`
	Gallery       []string      `json:"gallery"`
	QuantityLeft  float64       `json:"quantityLeft"`
	InStock       bool          `json:"inStock"`
	Price         float64       `json:"price"`
	Subtotal      float64       `json:"subtotal"`
	Total         float64       `json:"total"`
	Reviews       []*Reviews    `json:"reviews"`
}

// Products are a slice of Product
type Products []Product
