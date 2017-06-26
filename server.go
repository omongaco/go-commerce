package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iamclaytonray/go-commerce/controllers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())
	pc := controllers.NewProductController(getSession())

	/* API routes/endpoints */

	// Users
	r.POST("/api/v1/users", uc.CreateUser)

	// Products
	r.POST("/api/v1/products", pc.CreateProduct)

	fmt.Println("Server up")
	log.Fatal(http.ListenAndServe(":3000", r))

}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
