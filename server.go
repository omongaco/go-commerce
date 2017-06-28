package main

import (
	"log"
	"net/http"

	"github.com/iamclaytonray/go-commerce/controllers"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	// Create a new router
	r := httprouter.New()

	// Create a new CORS configuration
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
	})

	// Store the router in our CORS handler
	h := c.Handler(r)

	// Grab our controllers and get the MongoDB session
	uc := controllers.NewUserController(getSession())
	pc := controllers.NewProductController(getSession())

	/* API routes/endpoints */

	// User API Endpoints
	r.GET("/api/v1/users", uc.GetUsers)
	r.GET("/api/v1/users/:username", uc.GetUser)
	r.POST("/api/v1/users", uc.CreateUser)
	r.PUT("/api/v1/users/:username", uc.UpdateUser)
	r.DELETE("/api/v1/users/:username", uc.DeleteUser)

	// Product API Endpoints
	r.GET("/api/v1/products", pc.GetProducts)
	r.GET("/api/v1/products/:slug", pc.GetProduct)
	r.POST("/api/v1/products", pc.CreateProduct)
	r.PUT("/api/v1/products/:slug", pc.UpdateProduct)
	r.DELETE("/api/v1/products/:slug", pc.DeleteProduct)

	// Print that the server is listening and start the server on :3000
	log.Print("Server listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", logger(h)))

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
