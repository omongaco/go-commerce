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
	// Create a new router
	r := httprouter.New()

	// Create a new CORS configuration
	// c := cors.New(cors.Options{
	// 	AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	// 	AllowedOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
	// })

	// Store the router in our CORS handler
	// h := c.Handler(r)

	// Grab our controllers and get the MongoDB session
	orders := controllers.NewOrderController(getSession())
	users := controllers.NewUserController(getSession())
	products := controllers.NewProductController(getSession())

	/* API routes/endpoints */

	// Order API Endpoints
	r.GET("/api/v1/orders", orders.GetOrders)
	r.GET("/api/v1/orders/:slug", orders.GetOrder)
	r.POST("/api/v1/orders", orders.CreateOrder)
	r.PUT("/api/v1/orders/:slug", orders.UpdateOrder)
	// no delete?

	// User API Endpoints
	r.GET("/api/v1/users", users.GetUsers)
	r.GET("/api/v1/users/:username", users.GetUser)
	r.POST("/api/v1/users", users.CreateUser)
	r.PUT("/api/v1/users/:username", users.UpdateUser)
	r.DELETE("/api/v1/users/:username", users.DeleteUser)

	// Product API Endpoints
	r.GET("/api/v1/products", products.GetProducts)
	r.GET("/api/v1/products/:slug", products.GetProduct)
	r.POST("/api/v1/products", products.CreateProduct)
	r.PUT("/api/v1/products/:slug", products.UpdateProduct)
	r.DELETE("/api/v1/products/:slug", products.DeleteProduct)

	// Print that the server is listening and start the server on :3000
	fmt.Printf("Server listening on :3000\n")
	log.Fatal(http.ListenAndServe(":3000", logger(r)))

}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our MongoDB URI
	s, err := mgo.Dial("localhost")

	// Can a successful connection be made to Mongo? If not, panic at the disco :p
	if err != nil {
		panic(err)
	}

	// Return the session
	return s
}

func logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
