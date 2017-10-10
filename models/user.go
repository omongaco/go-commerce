package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User models our data for the database
type User struct {
	ID               bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt        time.Time     `json:"createdAt"`
	UpdatedAt        time.Time     `json:"updatedAt"`
	FirstName        string        `json:"firstName"`
	LastName         string        `json:"lastName"`
	Email            string        `json:"email"`
	Username         string        `json:"username"`
	Password         string        `json:"password"`
	Phone            string        `json:"phone"`
	Gender           string        `json:"gender"`
	Address          string        `json:"address"`
	City             string        `json:"city"`
	State            string        `json:"state"`
	Zip              string        `json:"zip"`
	Country          string        `json:"country"`
	FavoriteProducts []*Products   `json:"favoriteProducts"`
	BoughtProducts   []*Products   `json:"boughtProducts"`
	CartItems        []*Products   `json:"cartItems"`
}

// Users are a slice of User
type Users []User
