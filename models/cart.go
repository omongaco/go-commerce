package models

// Cart models our data for the database
type Cart struct {
	Products *Products `json:"products" bson:"products"`
}
